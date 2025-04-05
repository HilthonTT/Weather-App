package main

import (
	"net/http"

	"github.com/hilthontt/weather/internal/utils"
)

// healthcheckHandler godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck endpoint
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	string	"ok"
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.Env,
		"version": version,
	}

	if err := utils.JsonResponse(w, http.StatusOK, data); err != nil {
		utils.InternalServerError(w, r, err, app.logger)
	}
}
