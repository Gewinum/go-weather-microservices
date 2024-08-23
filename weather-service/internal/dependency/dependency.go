package dependency

import (
	amqprpc "github.com/0x4b53/amqp-rpc"
	"github.com/Gewinum/go-weather-microservices/weather-service/internal/config"
	"github.com/Gewinum/go-weather-microservices/weather-service/internal/server"
	"github.com/Gewinum/go-weather-microservices/weather-service/internal/services"
)

type Dependency struct {
	Config *config.Config

	WsService *services.WeatherService

	RpcServer       *amqprpc.Server
	WsServerHandler *server.WeatherServerHandler
}

func NewDependency() *Dependency {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	wsService := services.NewWeatherService()

	rpcServer := server.NewRPCServer(cfg.RabbitMQ)
	wsServerHandler := server.NewWeatherServerHandler(wsService, rpcServer)
	return &Dependency{
		Config: &cfg,

		WsService: wsService,

		RpcServer:       rpcServer,
		WsServerHandler: wsServerHandler,
	}
}
