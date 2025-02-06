package internal

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetWeatherByCity(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder("GET", "http://api.weatherapi.com/v1/current.json?q=Curitiba",
		httpmock.NewStringResponder(200, `{
			"location": {
				"name": "Curitiba",
				"region": "Parana",
				"country": "Brazil"
			},
			"current": {
				"temp_c": 20,
				"condition": {
					"text": "Partly cloudy"
				}
			}
		}`))

	// Test case
	t.Run("Valid city", func(t *testing.T) {
		weather, err := GetWeatherByCity("Curitiba")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if weather.Location.Name != "Curitiba" {
			t.Errorf("expected city Curitiba, got %v", weather.Location.Name)
		}
	})
	// Test case
	t.Run("Correct temp", func(t *testing.T) {
		weather, err := GetWeatherByCity("Curitiba")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if weather.Current.TempC != 20 {
			t.Errorf("expected temperature 20, got %v", weather.Current.TempC)
		}
	})
}
