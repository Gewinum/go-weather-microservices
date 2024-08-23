package server

import (
	amqprpc "github.com/0x4b53/amqp-rpc"
	"github.com/Gewinum/go-weather-microservices/common/models/weather"
	"github.com/Gewinum/go-weather-microservices/weather-service/internal/services"
)

type WeatherServerHandler struct {
	wsService *services.WeatherService
}

func NewWeatherServerHandler(wsService *services.WeatherService, connection *amqprpc.Server) *WeatherServerHandler {
	inst := &WeatherServerHandler{wsService: wsService}
	BindHandler[weather.RequestWeatherInformation, weather.ResponseWeatherInformation](connection, "weather_request", inst.requestWeather)
	return inst
}

func (hnd *WeatherServerHandler) requestWeather(command weather.RequestWeatherInformation) (weather.ResponseWeatherInformation, error) {
	return hnd.wsService.GetCurrentWeather(command.Lat, command.Lon)
}
