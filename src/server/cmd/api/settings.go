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

// GetSettings godoc
//
//	@Summary		Fetches user settings
//	@Description	Retrieves the settings for a user by their ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	store.Settings
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/{id}/settings [get]
func (app *application) getSettingsHandler(w http.ResponseWriter, r *http.Request) {
	settings := getSettingsFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, settings); err != nil {
		app.internalServerError(w, r, err)
	}
}

type UpdateSettingsPayload struct {
	TempFormat  store.TempFormat  `json:"temp_format" validate:"omitempty,oneof=celsius fahrenheit"`
	TimeFormat  store.TimeFormat  `json:"time_format" validate:"omitempty,oneof=24h 12h"`
	SpeedFormat store.SpeedFormat `json:"speed_format" validate:"omitempty,oneof=kmph mph"`
}

// UpdateSettings godoc
//
//	@Summary		Update user settings
//	@Description	Partially updates the settings for a user.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"User ID"
//	@Param			payload	body		UpdateSettingsPayload	true	"Settings to update"
//	@Success		200		{object}	store.Settings
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/{id}/settings [patch]
func (app *application) updateSettings(w http.ResponseWriter, r *http.Request) {
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

	settings.Update(payload.TempFormat, payload.TimeFormat, payload.SpeedFormat)

	if settings == nil {
		app.notFoundResponse(w, r, errors.New("Settings not found"))
		return
	}

	ctx := r.Context()

	err := app.store.Settings.Update(ctx, settings)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, settings); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) settingsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIDParam := chi.URLParam(r, "userID")
		userID, err := strconv.ParseInt(userIDParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		settings, err := app.store.Settings.GetByUserID(ctx, userID)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				// Create default settings
				settings = &store.Settings{
					UserID:      userID,
					TempFormat:  store.Celsius,
					TimeFormat:  store.TwentyFourHour,
					SpeedFormat: store.Kmph,
				}

				err := app.store.Settings.Create(ctx, settings)
				if err != nil {
					app.internalServerError(w, r, err)
					return
				}
			} else {
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
