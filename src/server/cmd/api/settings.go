package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hilthontt/weather/internal/store"
)

type settingsKey string

const settingsCtx settingsKey = "settings"

// adminRole is the role name (see migration 000003) allowed to access other
// users' settings.
const adminRole = "admin"

// GetSettings godoc
//
//	@Summary		Fetches user settings
//	@Description	Retrieves the settings for a user by their ID. Users may only access their own settings (admins may access any).
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	store.Settings
//	@Failure		400	{object}	error
//	@Failure		403	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/{id}/settings [get]
func (app *application) getSettingsHandler(w http.ResponseWriter, r *http.Request) {
	settings := getSettingsFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, settings); err != nil {
		app.internalServerError(w, r, err)
	}
}

// UpdateSettingsPayload is a partial update: omitted fields keep their
// current value.
type UpdateSettingsPayload struct {
	TempFormat  store.TempFormat  `json:"temp_format" validate:"omitempty,oneof=celsius fahrenheit"`
	TimeFormat  store.TimeFormat  `json:"time_format" validate:"omitempty,oneof=24h 12h"`
	SpeedFormat store.SpeedFormat `json:"speed_format" validate:"omitempty,oneof=kmph mph"`
}

// UpdateSettings godoc
//
//	@Summary		Update user settings
//	@Description	Partially updates the settings for a user. Omitted fields are left unchanged.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"User ID"
//	@Param			payload	body		UpdateSettingsPayload	true	"Settings to update"
//	@Success		200		{object}	store.Settings
//	@Failure		400		{object}	error
//	@Failure		403		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/{id}/settings [patch]
func (app *application) updateSettingsHandler(w http.ResponseWriter, r *http.Request) {
	var payload UpdateSettingsPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	settings := getSettingsFromCtx(r)
	if settings == nil {
		app.notFoundResponse(w, r, errors.New("settings not found"))
		return
	}

	settings.Update(payload.TempFormat, payload.TimeFormat, payload.SpeedFormat)

	if err := app.store.Settings.Update(r.Context(), settings); err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundResponse(w, r, err)
			return
		}

		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, settings); err != nil {
		app.internalServerError(w, r, err)
	}
}

// settingsContextMiddleware authorizes access to /users/{userID}/settings and
// loads (or lazily creates default) settings into the request context.
// It must run after AuthTokenMiddleware.
func (app *application) settingsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, errors.New("invalid user id"))
			return
		}

		currentUser := getUserFromContext(r)
		if currentUser == nil || (currentUser.ID != userID && currentUser.Role.Name != adminRole) {
			app.forbiddenResponse(w, r)
			return
		}

		ctx := r.Context()

		settings, err := app.store.Settings.GetByUserID(ctx, userID)
		if err != nil {
			if !errors.Is(err, store.ErrNotFound) {
				app.internalServerError(w, r, err)
				return
			}

			settings = store.DefaultSettings(userID)
			if err := app.store.Settings.Create(ctx, settings); err != nil {
				app.internalServerError(w, r, err)
				return
			}
		}

		ctx = context.WithValue(ctx, settingsCtx, settings)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getSettingsFromCtx(r *http.Request) *store.Settings {
	settings, _ := r.Context().Value(settingsCtx).(*store.Settings)

	return settings
}
