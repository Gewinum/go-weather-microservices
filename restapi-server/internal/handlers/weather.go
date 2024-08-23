package handlers

import (
	"github.com/Gewinum/go-weather-microservices/common/models/weather"
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/client"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WeatherHandler struct {
	wc *client.WeatherClient
}

func NewWeatherHandler(wc *client.WeatherClient) *WeatherHandler {
	return &WeatherHandler{wc: wc}
}

func (h *WeatherHandler) HandleForecast(c *gin.Context) {
	latStr, found := c.GetQuery("lat")
	if !found {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	lonStr, found := c.GetQuery("lon")
	if !found {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	info, err := h.wc.RequestWeatherInformation(weather.RequestWeatherInformation{Lat: lat, Lon: lon})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, info)
}
