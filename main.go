package main

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	api "roadmaps/projects/weather-api/internal/api"
	"roadmaps/projects/weather-api/internal/repository"
)

func main() {

	// Initialize Redis client
	redisAddr := "redis:6379" // Replace with your Redis server address
	repository.InitRedis(redisAddr)

	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/weather", api.GetWeather).Methods("GET")

	port := "8080"
	slog.Info("Server is starting", "port", port)
	// Start the HTTP server
	if err := http.ListenAndServe(":"+port, r); err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
