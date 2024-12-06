package models

type DayInfo struct {
	Datetime    string  `json:"datetime"`
	TempMax     float64 `json:"tempmax"`
	TempMin     float64 `json:"tempmin"`
	Temperature float64 `json:"temp"`
	Humidity    float64 `json:"humidity"`
	Sunset      string  `json:"sunset"`
	Sunrise     string  `json:"sunrise"`
}

// WeatherData represents the simplified weather data.
type WeatherData struct {
	//Temperature string `json:"temperature"`
	Timezone string `json:"timezone"`
	Address  string `json:"address"`
	//Temperature float64 `json:"temperature"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	Description string    `json:"description"`
	Days        []DayInfo `json:"days"`
}
