package dependency

import (
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/client"
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/config"
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Dependency struct {
	Config *config.Config

	WeatherCl *client.WeatherClient

	GinEngine *gin.Engine
}

func NewDependency() *Dependency {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	weatherClient := client.NewWeatherClient(cfg.RabbitMQ)

	weatherHandler := handlers.NewWeatherHandler(weatherClient)
	prometheusHandler := handlers.NewPrometheusHandler()
	e := gin.Default()
	e.Use(prometheusHandler.RPSMiddleware)
	e.Handle("GET", "/metrics", prometheusHandler.HandleMetrics)
	e.Handle("GET", "/", weatherHandler.HandleForecast)

	return &Dependency{
		Config: &cfg,

		WeatherCl: weatherClient,

		GinEngine: e,
	}
}
