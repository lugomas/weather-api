package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"roadmaps/projects/weather-api/internal/models"
	"roadmaps/projects/weather-api/internal/services"
)

// GetWeather handles the request for weather data.
func GetWeather(w http.ResponseWriter, r *http.Request) {
	// Parse the `address` query parameter.
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}

	// Fetch weather data.
	weatherData, err := services.FetchWeatherData(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert map to JSON
	jsonData, err := json.Marshal(weatherData)
	if err != nil {
		log.Printf("ERROR: Error marshalling: %v", err)
		fmt.Fprintf(w, "Error marshalling: %v", err)
		return
	}
	// Unmarshal JSON to struct
	var weather models.WeatherData
	err = json.Unmarshal(jsonData, &weather)
	if err != nil {
		log.Printf("ERROR: Error unmarshalling: %v", err)
		fmt.Fprintf(w, "Error unmarshalling: %v", err)
		return
	}

	// Print the result
	fmt.Fprintf(w, "Address: %s \n", weather.Address)
	fmt.Fprintf(w, "Timezone: %s \n", weather.Timezone)
	fmt.Fprintf(w, "Latitude: %f \n", weather.Latitude)
	fmt.Fprintf(w, "Longitude: %f \n", weather.Longitude)
	fmt.Fprintf(w, "Description: %s \n", weather.Description)
	//fmt.Fprintf(w, "Datetime: %s \n", weather.Days.Datetime)
	//fmt.Fprintf(w, "TempMax: %s \n", weather.Days.TempMax)
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

	// Return the weather data as JSON.
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(weatherData)

	//// Extract the "timezone" field
	//timezone, _ := weatherData["timezone"].(string)
	//address, _ = weatherData["address"].(string)
	//// Return the weather Data
	//fmt.Fprintf(w, "address: %s \n", address)
	//fmt.Fprintf(w, "timezone: %s \n", timezone)
	//fmt.Fprintf(w, "----------------------------------------------- \n")

}
