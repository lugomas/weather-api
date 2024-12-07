package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"roadmaps/projects/weather-api/internal/models"
	"roadmaps/projects/weather-api/internal/services"
)

// GetWeather handles the request for weather data.
func GetWeather(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default()

	// Parse the `address` query parameter.
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		logger.Warn("Missing address query parameter")
		return
	}

	// Fetch weather data.
	weatherData, err := services.FetchWeatherData(address)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		logger.Error("Failed to fetch weather data", "error", err, "address", address)
		return
	}

	// Convert map to JSON
	jsonData, err := json.Marshal(weatherData)
	if err != nil {
		http.Error(w, "Error marshalling weather data", http.StatusInternalServerError)
		logger.Error("Error marshalling weather data", "error", err)
		return
	}

	// Unmarshal JSON to struct
	var weather models.WeatherData
	err = json.Unmarshal(jsonData, &weather)
	if err != nil {
		http.Error(w, "Error unmarshalling weather data", http.StatusInternalServerError)
		logger.Error("Error unmarshalling weather data", "error", err)
		return
	}

	// Log successful weather data retrieval
	logger.Info("Weather data retrieved successfully", "address", weather.Address, "timezone", weather.Timezone)

	// Print the result
	fmt.Fprintf(w, "Address: %s \n", weather.Address)
	fmt.Fprintf(w, "Timezone: %s \n", weather.Timezone)
	fmt.Fprintf(w, "Latitude: %f \n", weather.Latitude)
	fmt.Fprintf(w, "Longitude: %f \n", weather.Longitude)
	fmt.Fprintf(w, "Description: %s \n", weather.Description)
	fmt.Fprintf(w, "----------------------------------------------- \n")

	// Access datetime and tempmin
	for _, day := range weather.Days {
		fmt.Fprintf(w, "Datetime: %s \n", day.Datetime)
		fmt.Fprintf(w, "TempMin: %.1f\n", day.TempMin)
		fmt.Fprintf(w, "TempMax: %.1f\n", day.TempMax)
		fmt.Fprintf(w, "Temperature: %.1f\n", day.Temperature)
		fmt.Fprintf(w, "Sunset: %s\n", day.Sunset)
		fmt.Fprintf(w, "Sunrise: %s\n", day.Sunrise)
		fmt.Fprintf(w, "Humidity: %.1f\n", day.Humidity)
	}
	fmt.Fprintf(w, "----------------------------------------------- \n")
}
