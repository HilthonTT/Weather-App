package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hilthontt/weather/internal/utils"
	"go.uber.org/zap"
)

type userKey string

const UserCtx userKey = "user"

type Handler struct {
	logger *zap.SugaredLogger
}

func NewHandler(logger *zap.SugaredLogger) *Handler {
	return &Handler{logger}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/users/register", h.registerUserHandler)
	r.Post("/users/login", h.loginUserHandler)
}

func (h *Handler) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, http.StatusOK, "Ok!")
}

func (h *Handler) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, http.StatusOK, "Ok!")
}
