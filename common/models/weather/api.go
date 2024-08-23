package weather

type RequestWeatherInformation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type ResponseWeatherInformation struct {
	WeatherCode   float64 `json:"weather_code"`
	Temperature   float64 `json:"temperature"`
	Time          int64   `json:"time"`
	WindDirection float64 `json:"wind_direction"`
	WindSpeed     float64 `json:"wind_speed"`
}
