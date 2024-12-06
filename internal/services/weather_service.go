package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"roadmaps/projects/weather-api/internal/repository"
)

// FetchWeatherData retrieves weather data for a given address.
func FetchWeatherData(address string) (map[string]interface{}, error) {
	// Check Redis cache for data
	cachedData, err := repository.GetCachedWeather(address)
	if err != nil {
		log.Printf("ERROR: error checking cache for address %s: %v", address, err)
	} else if cachedData != "" {
		// Cache hit
		log.Printf("INFO: data retrieved from cache for address: %s", address)
		log.Printf("INFO: API response: %v", string(cachedData))
		var cachedResult map[string]interface{}
		if err := json.Unmarshal([]byte(cachedData), &cachedResult); err != nil {
			log.Printf("ERROR: failed to unmarshal cached data: %v", err)
		} else {
			return cachedResult, nil
		}
	} else {
		log.Printf("INFO: cache miss for address: %s", address)
	}

	// Cache miss - Fetch from external API
	log.Printf("INFO: fetching data from 3rd party API for address: %s", address)
	apiKey := "<include your API Key>" // Replace with your actual API key.
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/today?unitGroup=metric&key=%s&contentType=json", address, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ERROR: failed to call weather API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ERROR: weather API returned status %d", resp.StatusCode)
	}

	log.Printf("INFO: successfully fetched data from 3rd party API for address: %s", address)

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("ERROR: failed to decode weather API response: %w", err)
	}

	// Cache the fetched data in Redis
	dataBytes, err := json.Marshal(result)
	if err == nil {
		if cacheErr := repository.SetCachedWeather(address, string(dataBytes)); cacheErr != nil {
			log.Printf("ERROR: failed to cache data for address %s: %v", address, cacheErr)
		} else {
			log.Printf("INFO: data cached successfully for address: %s", address)
			log.Printf("INFO: API response: %v", string(dataBytes))
		}
	} else {
		log.Printf("ERROR: failed to marshal data for caching: %v", err)
	}
	return result, nil
}
