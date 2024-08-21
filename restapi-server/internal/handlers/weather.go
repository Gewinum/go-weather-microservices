package handlers

import (
	"github.com/Gewinum/go-weather-microservices/common/models/weather"
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WeatherHandler struct {
	wc *client.WeatherClient
}

func NewWeatherHandler(wc *client.WeatherClient) *WeatherHandler {
	return &WeatherHandler{wc: wc}
}

func (h *WeatherHandler) HandleForecast(c *gin.Context) {
	city, found := c.GetQuery("city")
	if !found {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	info, err := h.wc.RequestWeatherInformation(weather.RequestWeatherInformation{City: city})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, info)
}
