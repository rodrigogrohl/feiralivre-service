package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/config"
	"github.com/sirupsen/logrus"
)

var (
	servicesCounter *prometheus.CounterVec
	// marketIndicatorGauge *prometheus.GaugeVec
)

func Expose() error {
	http.Handle("/metrics", promhttp.Handler())
	logrus.WithField("port", config.PrometheusPort).Info("initializing /metrics endpoint")
	return http.ListenAndServe(fmt.Sprintf(":%d", config.PrometheusPort), nil)
}

func InitializeMonitors() {
	servicesCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "streetmarket",
			ConstLabels: prometheus.Labels{
				"env": config.Environment,
			},
			Help: "Service API Calls",
		},
		[]string{"service", "method", "status"},
	)
}

func AddCounter(service, method, status string) {
	if servicesCounter == nil {
		logrus.Info("metrics not initialized")
		return
	}
	servicesCounter.WithLabelValues(service, method, status).Inc()
}

func AddCounterErr(service, method string, err error) {
	status := "OK"
	if err != nil {
		status = err.Error()
	}
	AddCounter(service, method, status)
}
