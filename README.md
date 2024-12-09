# Weather API App
A Go-based Weather API application that allows users to fetch real-time weather data for specific addresses. The application features Redis integration for efficient caching and enhanced performance.

## Features
- Weather Data Retrieval: Fetch Weather data from a 3rd party API for any address. Powered by [Visual Crossing’s API](https://www.visualcrossing.com/weather-api)
- Caching with Redis: Improves response time and reduces external API calls.
- Scalable Deployment: Easily scalable and containerized for seamless deployment using Docker.
- Simple Go Server Setup: Easy-to-understand and lightweight server implementation using Go.
- Port Configuration: Runs on port 8080 by default, configurable for flexibility.

## Prerequisites
Before running the Weather API application, ensure you have the following installed on your machine:
- **Go**: Version 1.22.3 (darwin/arm64)
- **Docker**: Make sure Docker is installed and running.
- **Docker Compose**: This is included with Docker Desktop.
- **cURL**: Ensure cURL is installed for testing API requests.


## Configuration
- Obtain an API key from [Visual Crossing’s API](https://www.visualcrossing.com/weather-api).
- Add your API key to the environment variable *WEATHER_API_KEY*  
  ```export WEATHER_API_KEY=your_api_key_here```

## Installation
To install and run the `weather-api` app locally, clone the repository and build the Go binary, following the steps below:
```
git clone https://github.com/lugomas/weather-api.git
cd weather-api
go build -o weather-api
```

## Running the Application
1. Start the app:  
   ```docker-compose up --build```
 
## Usage
1. Access the app:
   Open your web browser and navigate to:  
   ```http://localhost:8080/weather?address=sao%20paulo```  
   Or open a new terminal and test it via cURL :  
   ```curl -v "http://localhost:8080/weather?address=london"```

## License
This project is licensed under the MIT License.

## Project Inspiration
This project was developed based on the guidelines provided by [roadmap.sh's Weather API project](https://roadmap.sh/projects/weather-api-wrapper-service)
