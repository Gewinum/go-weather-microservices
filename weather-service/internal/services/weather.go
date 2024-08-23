package services

import (
	"context"
	"github.com/Gewinum/go-weather-microservices/common/models/weather"
	"github.com/hectormalot/omgo"
)

type WeatherService struct {
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (w *WeatherService) GetCurrentWeather(lat float64, lon float64) (weather.ResponseWeatherInformation, error) {
	c, _ := omgo.NewClient()
	loc, _ := omgo.NewLocation(lat, lon)
	res, err := c.CurrentWeather(context.Background(), loc, nil)
	if err != nil {
		return weather.ResponseWeatherInformation{}, err
	}
	return weather.ResponseWeatherInformation{
		WeatherCode:   res.WeatherCode,
		Temperature:   res.Temperature,
		Time:          res.Time.Unix(),
		WindDirection: res.WindDirection,
		WindSpeed:     res.WindSpeed,
	}, nil
}
