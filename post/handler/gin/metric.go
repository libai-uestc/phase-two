package handler

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SERVICE = "post"

var (
	requestCounter = promauto.NewCounterVec(prometheus.CounterOpts{Name: "request_counter"}, []string{"service", "interface"})
	requestTimer   = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "request_timer"}, []string{"service", "interface"})
)

func Metric(ctx *gin.Context) {
	begin := time.Now()
	ctx.Next()
	ifc := mappingUrl(ctx)
	requestCounter.WithLabelValues(SERVICE, ifc).Inc()
	requestTimer.WithLabelValues(SERVICE, ifc).Set(float64(time.Since(begin).Milliseconds()))
}

var (
	restfulMapping = map[string]string{"id": ":id"}
)

func mappingUrl(ctx *gin.Context) string {
	url := ctx.Request.URL.Path
	for _, p := range ctx.Params {
		if value, exists := restfulMapping[p.Key]; exists {
			url = strings.Replace(url, p.Value, value, 1)
		}
	}
	return url
}
