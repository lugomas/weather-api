package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	api "roadmaps/projects/weather-api/internal/api"
	"roadmaps/projects/weather-api/internal/repository"
)

func main() {

	// Initialize Redis client
	redisAddr := "localhost:6379" // Replace with your Redis server address
	repository.InitRedis(redisAddr)

	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/weather", api.GetWeather).Methods("GET")

	port := "8080"
	log.Printf("INFO: server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
