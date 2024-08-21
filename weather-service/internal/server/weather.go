package server

import (
	amqprpc "github.com/0x4b53/amqp-rpc"
	"github.com/Gewinum/go-weather-microservices/common/models/weather"
	"math/rand"
)

func RegisterWeatherCommands(connection *amqprpc.Server) {
	BindHandler[weather.RequestWeatherInformation, weather.ResponseWeatherInformation](connection, "weather_request", requestWeather)
}

func requestWeather(command weather.RequestWeatherInformation) (weather.ResponseWeatherInformation, error) {
	return weather.ResponseWeatherInformation{
		Temperature: rand.Float64(),
		Humidity:    rand.Float64(),
	}, nil
}
