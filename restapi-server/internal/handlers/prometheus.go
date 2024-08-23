package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusHandler struct {
	GinGauge prometheus.Gauge
}

func NewPrometheusHandler() *PrometheusHandler {
	inst := &PrometheusHandler{}
	inst.GinGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gin_total_requests",
		Help: "Total requests amount in Gin",
	})
	prometheus.MustRegister(inst.GinGauge)
	return inst
}

func (p *PrometheusHandler) RPSMiddleware(c *gin.Context) {
	p.GinGauge.Add(1.0)
}

func (p *PrometheusHandler) HandleMetrics(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}
