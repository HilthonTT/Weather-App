package users

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hilthontt/weather/internal/auth"
	"github.com/hilthontt/weather/internal/config"
	"github.com/hilthontt/weather/internal/db"
	"github.com/hilthontt/weather/internal/utils"
	"go.uber.org/zap"
)

type Handler struct {
	logger        *zap.SugaredLogger
	store         *UserStore
	authConfig    config.AuthConfig
	authenticator auth.Authenticator
}

func NewHandler(
	logger *zap.SugaredLogger,
	store *UserStore,
	authConfig config.AuthConfig,
	authenticator auth.Authenticator) *Handler {
	return &Handler{logger, store, authConfig, authenticator}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/users/register", h.registerUserHandler)
	r.Post("/users/login", h.loginUserHandler)
}

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

// registerUserHandler godoc
//
//	@Summary		Registers a new user
//	@Description	Creates a new user account with a username, email, and password
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterUserPayload	true	"New user registration data"
//	@Success		201		{object}	User				"Created user object"
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/users/register [post]
func (h *Handler) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.BadRequestResponse(w, r, err, h.logger)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.BadRequestResponse(w, r, err, h.logger)
		return
	}

	user := &User{
		Username: payload.Username,
		Email:    payload.Email,
		Role: Role{
			Name: "user",
		},
	}

	// hash the user password
	if err := user.Password.Set(payload.Password); err != nil {
		utils.InternalServerError(w, r, err, h.logger)
		return
	}

	if err := utils.JsonResponse(w, http.StatusCreated, user); err != nil {
		utils.InternalServerError(w, r, err, h.logger)
	}
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

// loginUserHandler godoc
//
//	@Summary		Logs in a user
//	@Description	Authenticates a user with email and password, and returns a JWT token upon successful login
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		LoginUserPayload	true	"User login credentials"
//	@Success		201		{string}	string				"JWT token"
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/users/login [post]
func (h *Handler) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload LoginUserPayload
	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.BadRequestResponse(w, r, err, h.logger)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.BadRequestResponse(w, r, err, h.logger)
		return
	}

	user, err := h.store.GetByEmail(r.Context(), payload.Email)

	if err != nil {
		switch err {
		case db.ErrNotFound:
			utils.UnauthorizedErrorResponse(w, r, err, h.logger)
		default:
			utils.InternalServerError(w, r, err, h.logger)
		}
		return
	}

	if err := user.Password.Compare(payload.Password); err != nil {
		utils.UnauthorizedErrorResponse(w, r, err, h.logger)
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(h.authConfig.Token.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": h.authConfig.Token.Iss,
		"aud": h.authConfig.Token.Iss,
	}

	token, err := h.authenticator.GenerateToken(claims)
	if err != nil {
		utils.InternalServerError(w, r, err, h.logger)
		return
	}

	if err := utils.JsonResponse(w, http.StatusCreated, token); err != nil {
		utils.InternalServerError(w, r, err, h.logger)
	}
}
