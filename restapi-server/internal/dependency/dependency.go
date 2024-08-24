package dependency

import (
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/client"
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/config"
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/handlers"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
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

	store := ratelimit.RedisStore(&ratelimit.RedisOptions{
		RedisClient: redis.NewClient(&redis.Options{
			Addr: "redis:6379",
		}),
		Rate:  time.Second,
		Limit: 1,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: rateLimitErrorHandler,
		KeyFunc:      rateLimitKeyFunc,
	})

	e := gin.Default()
	e.Use(prometheusHandler.RPSMiddleware)
	e.Handle("GET", "/metrics", prometheusHandler.HandleMetrics)
	e.Handle("GET", "/", mw, weatherHandler.HandleForecast)

	return &Dependency{
		Config: &cfg,

		WeatherCl: weatherClient,

		GinEngine: e,
	}
}

func rateLimitKeyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func rateLimitErrorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}
