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
		return
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
