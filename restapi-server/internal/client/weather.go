package client

import (
	amqprpc "github.com/0x4b53/amqp-rpc"
	"github.com/Gewinum/go-weather-microservices/common/models/weather"
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/config"
)

func NewWeatherClient(rabbitMq config.RabbitMQConfig) *WeatherClient {
	connection := amqprpc.NewClient(BuildRabbitMQURL(rabbitMq))
	return &WeatherClient{
		connection: connection,
	}
}

type WeatherClient struct {
	connection *amqprpc.Client
}

func (wc *WeatherClient) RequestWeatherInformation(command weather.RequestWeatherInformation) (*weather.ResponseWeatherInformation, error) {
	return SendCommand[weather.RequestWeatherInformation, weather.ResponseWeatherInformation](wc.connection, "weather_request", command)
}
