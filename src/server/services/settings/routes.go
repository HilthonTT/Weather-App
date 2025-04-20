package settings

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hilthontt/weather/internal/utils"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.SugaredLogger
	store  *SettingsStore
}

func NewHandler(logger *zap.SugaredLogger, store *SettingsStore) *Handler {
	return &Handler{logger, store}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/settings", func(r chi.Router) {
		r.Get("/", h.checkSettingsOwnership(h.getSettings))
	})
}

// GetSettings godoc
//
//	@Summary		Retrieve user settings
//	@Description	Fetches the settings associated with the authenticated user
//	@Security		BearerAuth
//	@Tags			settings
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Settings
//	@Failure		500	{object}	error
//	@Router			/settings [get]
func (h *Handler) getSettings(w http.ResponseWriter, r *http.Request) {
	settings := GetSettingsFromContext(r)

	if err := utils.JsonResponse(w, http.StatusOK, settings); err != nil {
		utils.InternalServerError(w, r, err, h.logger)
	}
}
