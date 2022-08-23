package main

import (
	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/config"
	"github.com/rodrigogrohl/feiralivre-service/internal/presentation/metrics"
	"github.com/rodrigogrohl/feiralivre-service/internal/presentation/rest"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("===== config loader =====")
	config.Initialize("prod")

	logrus.Info("===== initializing metrics =====")
	metrics.InitializeMonitors()
	go metrics.Expose()

	logrus.Info("===== initializing streetmarket application =====")
	rest.InitRestService()
}
