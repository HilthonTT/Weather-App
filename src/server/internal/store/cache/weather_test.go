package cache

import "testing"

func TestCoordKey_RoundsSoRequestAndResponseMatch(t *testing.T) {
	// A raw GPS fix from the app vs. the coordinates OpenWeather echoes back.
	request := coordKey("weather", 51.483312, -0.083341)
	response := coordKey("weather", 51.4833, -0.0833)

	if request != response {
		t.Errorf("keys differ: %q vs %q", request, response)
	}
}

func TestCityKey_IsCaseInsensitive(t *testing.T) {
	if cityKey("weather", " london ") != cityKey("weather", "London") {
		t.Error("city keys should be normalized")
	}
}
