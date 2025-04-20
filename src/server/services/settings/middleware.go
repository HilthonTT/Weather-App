package settings

import (
	"net/http"

	"github.com/hilthontt/weather/internal/utils"
	"github.com/hilthontt/weather/services/users"
)

func (h Handler) checkSettingsOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := users.GetUserFromContext(r)
		settings := GetSettingsFromContext(r)

		if settings.UserID == user.ID {
			next.ServeHTTP(w, r)
			return
		}

		utils.ForbiddenResponse(w, r, h.logger)
	})
}
