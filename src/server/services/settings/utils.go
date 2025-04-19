package settings

import "net/http"

type settingsKey string

const SettingsCtx settingsKey = "settings"

func GetSettingsFromContext(r *http.Request) *Settings {
	settings, _ := r.Context().Value(SettingsCtx).(*Settings)

	return settings
}
