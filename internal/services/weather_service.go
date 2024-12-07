package services

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"roadmaps/projects/weather-api/internal/repository"
)

// FetchWeatherData retrieves weather data for a given address.
func FetchWeatherData(address string) (map[string]interface{}, error) {
	logger := slog.Default()

	// Check Redis cache for data
	cachedData, err := repository.GetCachedWeather(address)
	if err != nil {
		logger.Error("Error checking cache for address", "address", address, "error", err)
	} else if cachedData != "" {
		// Cache hit
		logger.Info("Data retrieved from cache", "address", address)
		logger.Info("api response: ", "response", cachedData)
		var cachedResult map[string]interface{}
		if err := json.Unmarshal([]byte(cachedData), &cachedResult); err != nil {
			logger.Error("Failed to unmarshal cached data", "address", address, "error", err)
		} else {
			return cachedResult, nil
		}
	} else {
		logger.Info("Cache miss", "address", address)
	}

	// Cache miss - Fetch from external API
	logger.Info("Fetching data from external API", "address", address)
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		slog.Error("API key not found in environment variables")
		return nil, fmt.Errorf("API key not configured")
	}

	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/today?unitGroup=metric&key=%s&contentType=json", address, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Failed to call weather API", "address", address, "error", err)
		return nil, fmt.Errorf("failed to call weather API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("Weather API returned non-OK status", "address", address, "status", resp.StatusCode)
		return nil, fmt.Errorf("weather API returned status %d", resp.StatusCode)
	}

	logger.Info("Successfully fetched data from external API", "address", address)

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Error("Failed to decode weather API response", "address", address, "error", err)
		return nil, fmt.Errorf("failed to decode weather API response: %w", err)
	}

	// Cache the fetched data in Redis
	dataBytes, err := json.Marshal(result)
	if err != nil {
		logger.Error("Failed to marshal data for caching", "address", address, "error", err)
	} else {
		if cacheErr := repository.SetCachedWeather(address, string(dataBytes)); cacheErr != nil {
			logger.Error("Failed to cache data", "address", address, "error", cacheErr)
		} else {
			logger.Info("Weather data cached successfully", "address", address, "expiration", "10m")
			logger.Info("api response: ", "response", string(dataBytes))
		}
	}

	return result, nil
}
