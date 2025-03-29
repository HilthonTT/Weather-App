package utils

import (
	"log"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	WriteJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func ForbiddenResponse(w http.ResponseWriter, r *http.Request) {
	//	app.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error")

	WriteJSONError(w, http.StatusForbidden, "forbidden")
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

func ConflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Errorf("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	WriteJSONError(w, http.StatusConflict, err.Error())
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Warnf("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	WriteJSONError(w, http.StatusNotFound, "not found")
}

func UnauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//app.logger.Warnf("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	WriteJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func UnauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	// app.logger.Warnf("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	WriteJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func RateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	log.Printf("rate limit exceeded for method: %v and path: %v", r.Method, r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	WriteJSONError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}
