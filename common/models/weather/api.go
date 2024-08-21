package weather

type RequestWeatherInformation struct {
	City string `json:"city"`
}

type ResponseWeatherInformation struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
