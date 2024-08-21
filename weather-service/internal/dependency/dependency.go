package dependency

import (
	amqprpc "github.com/0x4b53/amqp-rpc"
	"github.com/Gewinum/go-weather-microservices/weather-service/internal/config"
	"github.com/Gewinum/go-weather-microservices/weather-service/internal/server"
)

type Dependency struct {
	Config *config.Config

	RpcServer *amqprpc.Server
}

func NewDependency() *Dependency {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	rpcServer := server.NewRPCServer(cfg.RabbitMQ)
	server.RegisterWeatherCommands(rpcServer)
	return &Dependency{
		Config: &cfg,

		RpcServer: rpcServer,
	}
}
