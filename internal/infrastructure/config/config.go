package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Application string
	Environment string

	RestPort       int
	PrometheusPort int

	DbDriver     string
	DbConnection string

	LogToFile   bool
	LogFilename string

	BasePath string
)

func Initialize(filename string) {
	if filename == "" {
		filename = "prod"
	}

	viper.AutomaticEnv()

	viper.SetConfigName(filename)
	viper.SetConfigType("json")
	viper.AddConfigPath("/app/configs")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Panic("panic reading config file")
	}

	loadGlobals()
	ConfigLog()

	logrus.WithField("config_keys", viper.AllKeys()).Info()
	for _, k := range viper.AllKeys() {
		logrus.WithField("key", k).WithField("value", viper.Get(k)).Info()
	}
}

func ConfigLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "02-01-2006T15:04:05.000Z"})
	logrus.SetLevel(logrus.InfoLevel)
	// logrus.SetOutput(ioutil.Discard)
	// logrus.SetReportCaller(true)

	if LogToFile {
		f, err := os.OpenFile(BasePath+LogFilename, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			logrus.WithError(err).Panic("error creating log file")
		}
		logrus.SetOutput(f)
	}

}

func loadGlobals() {
	Application = requiredString("app")
	Environment = requiredString("environment")

	RestPort = requiredInt("rest_port")
	PrometheusPort = requiredInt("prometheus_port")

	DbDriver = requiredString("db.driver")
	DbConnection = requiredString("db.conn")

	LogToFile = viper.GetBool("log.to_file")
	LogFilename = viper.GetString("log.filename")

	BasePath = getBasePath()
}

func requiredString(key string) string {
	result := viper.GetString(key)
	if result == "" {
		logrus.WithField("key", key).Panic("string setting not found")
	}
	return result
}

func requiredInt(key string) int {
	result := viper.GetInt(key)
	if result == 0 {
		logrus.WithField("key", key).Panic("int setting not found")
	}
	return result
}

func getBasePath() string {
	dir, err := os.Getwd()
	if err != nil {
		logrus.Panic(err)
	}
	subs := strings.SplitN(dir, "feiralivre-service", 2)
	return subs[0] + "feiralivre-service/"
}
