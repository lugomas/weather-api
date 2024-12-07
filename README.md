# Weather API App
A weather API application built with Go, enabling users to retrieve weather data for specific addresses.
## Features
- Fetch Weather data from a 3rd party API. 
- Powered by [Visual Crossing’s API](https://www.visualcrossing.com/weather-api)

## Prerequisites
- Go installed on your machine.
- Redis installed and running.

## Installation
To install and run the `weather-api` app locally, clone the repository and build the Go binary, following the steps below:
```
git clone https://github.com/lugomas/weather-api.git
cd weather-api
go build -o weather-api
```

## Configuration
- Obtain an API key from [Visual Crossing’s API](https://www.visualcrossing.com/weather-api).
- Update the code to include your API key.

## Running the Application
1. Ensure Redis is installed and running:
    For macOS:
    Install Redis: 
    ```brew install redis```
    Start Redis: 
    ```brew services start redis```
2. Start the application: 
   ```./weather-api```
3. Access the app:
   Open your web browser and navigate to 
   ```http://localhost:8080/weather?address=sao%20paulo```
   Or test via cURL
   ```curl -v "http://localhost:8080/weather?address=london"```
 
## License
This project is licensed under the MIT License.

## Project Inspiration
This project was developed based on the guidelines provided by [roadmap.sh's Weather API project](https://roadmap.sh/projects/weather-api-wrapper-service)

## Project Backlog
- Create dockerfile for redis to run it as container
  - refactor code if necessary
- Handle cases such as:
  - Redis stopped working. Retrieve directly from 3rd party API
  - Redis didn't start. Retrieve directly from 3rd party API
  - APP couldn't connect to REDIS. Retrieve directly from 3rd party API
