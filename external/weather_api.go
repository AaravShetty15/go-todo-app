package external

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}

func GetWeather(lat, lon float64) (WeatherResponse, error) {

	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true",
		lat,
		lon,
	)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)

	return weather, err
}