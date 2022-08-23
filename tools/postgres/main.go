package main

import (
	"context"
	"net"
	"os"
	"os/exec"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/config"
	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/repository/database"
	"github.com/sirupsen/logrus"
)

const (
	targetPath = "./assets/FEIRAS_LIVRES/CSV/DEINFO_DADOS_AB_FEIRASLIVRES/"
	targetFile = "DEINFO_AB_FEIRASLIVRES_2014.csv"
	loadModel  = "./scripts/db/street_market.sql"
)

type StreetMarket struct {
	Id            int64   `csv:"ID"`
	Longitude     float64 `csv:"LONG"`
	Latitude      float64 `csv:"LAT"`
	SectorCense   int64   `csv:"SETCENS"`
	AreaPonderate int64   `csv:"AREAP"`
	DistrictCode  int64   `csv:"CODDIST"`
	District      string  `csv:"DISTRITO"`
	SubTownCode   int64   `csv:"CODSUBPREF"`
	SubTown       string  `csv:"SUBPREFE"`
	Region5       string  `csv:"REGIAO5"`
	Region8       string  `csv:"REGIAO8"`
	Name          string  `csv:"NOME_FEIRA"`
	Registry      string  `csv:"REGISTRO"`
	Address       string  `csv:"LOGRADOURO"`
	Number        string  `csv:"NUMERO"`
	Neighborhood  string  `csv:"BAIRRO"`
	Reference     string  `csv:"reference"`
}

func main() {
	config.InitializeTest()
	ctx := context.Background()

	go exec.CommandContext(ctx, "/usr/local/bin/docker-compose", "-f /home/rodrigo/go/src/github.com/rodrigogrohl/feiralivre-service/tools/db/docker-compose.yaml start").CombinedOutput()

	dbStruct := database.DatabaseConnect()
	if IsDatabaseRunning("127.0.0.1", "15432") {
		b, err := os.ReadFile(loadModel) // just pass the file name
		if err != nil {
			logrus.WithError(err).Panic()
		}

		sqlStatement := string(b)
		_, err = dbStruct.DB.Exec(sqlStatement)
		if err != nil {
			logrus.WithError(err).Panic()
		}
	}
}

func IsDatabaseRunning(host, port string) bool {
	for {
		logrus.Info("checking gcp proxy connection")
		if wasOpened(host, port) {
			logrus.Info("database running with success")
			return true
		}
		time.Sleep(1 * time.Second)
	}
}

func wasOpened(host string, port string) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		logrus.WithError(err).Debug("waiting for connection")
		return false
	}
	if conn != nil {
		defer conn.Close()
		logrus.WithField("opened", net.JoinHostPort(host, port)).Info()
		return true
	}
	return false
}
