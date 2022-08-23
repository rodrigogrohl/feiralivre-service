package main

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/config"
	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/repository"
	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/repository/database"
	"github.com/rodrigogrohl/feiralivre-service/pkg/canonical"
	"github.com/sirupsen/logrus"
)

const (
	targetPath = "assets/FEIRAS_LIVRES/CSV/DEINFO_DADOS_AB_FEIRASLIVRES/"
	targetFile = "DEINFO_AB_FEIRASLIVRES_2014.csv"
	loadModel  = "./scripts/db/street_market.sql"
)

// type StreetMarket struct {
// 	Id            int64   `csv:"ID"`
// 	Longitude     float64 `csv:"LONG"`
// 	Latitude      float64 `csv:"LAT"`
// 	SectorCense   int64   `csv:"SETCENS"`
// 	AreaPonderate int64   `csv:"AREAP"`
// 	DistrictCode  int64   `csv:"CODDIST"`
// 	District      string  `csv:"DISTRITO"`
// 	SubTownCode   int64   `csv:"CODSUBPREF"`
// 	SubTown       string  `csv:"SUBPREFE"`
// 	Region5       string  `csv:"REGIAO5"`
// 	Region8       string  `csv:"REGIAO8"`
// 	Name          string  `csv:"NOME_FEIRA"`
// 	Registry      string  `csv:"REGISTRO"`
// 	Address       string  `csv:"LOGRADOURO"`
// 	Number        string  `csv:"NUMERO"`
// 	Neighborhood  string  `csv:"BAIRRO"`
// 	Reference     string  `csv:"reference"`
// }

func main() {
	config.InitializeTest()

	// UNMARSHAL CSV FILE
	f, err := os.Open(config.BasePath + targetPath + targetFile)
	if err != nil {
		logrus.Fatal(err)
	}

	markets := []*canonical.StreetMarket{}

	reader := csv.NewReader(f)
	for {
		sm, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			if serr, ok := err.(*csv.ParseError); ok {
				logrus.WithField("street_market", sm).WithError(serr).Warn("skipping unmarshal due to missing fields. check me into CSV file!")
			} else {
				logrus.Fatal(err)
			}
		} else {
			if sm[0] == "ID" {
				continue
			}
			structResult := Unmarshal(sm)
			if structResult != nil {
				markets = append(markets, structResult)
			}
		}
	}

	// RELOAD DATABASE MODEL
	dbStruct := database.DatabaseConnect()
	err = database.ResetAllModel(dbStruct)
	if err != nil {
		logrus.WithError(err).Panic()
	}

	// LOAD DATA
	repos := repository.InitStreetMarketRepository()
	for _, market := range markets {
		repos.Insert(context.Background(), market)
	}

}

func ToInt64(val string) int64 {
	res, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		logrus.WithError(err).Panic()
	}
	return res
}

func ToFloat(val string) float64 {
	res, err := strconv.ParseFloat(val, 64)
	if err != nil {
		logrus.WithError(err).Panic()
	}
	return res
}

func Unmarshal(sm []string) *canonical.StreetMarket {
	result := &canonical.StreetMarket{
		Id:            ToInt64(sm[0]),
		Longitude:     ToFloat(sm[1]),
		Latitude:      ToFloat(sm[2]),
		SectorCense:   ToInt64(sm[3]),
		AreaPonderate: ToInt64(sm[4]),
		DistrictCode:  ToInt64(sm[5]),
		District:      sm[6],
		SubTownCode:   ToInt64(sm[7]),
		SubTown:       sm[8],
		Region5:       sm[9],
		Region8:       sm[10],
		Name:          sm[11],
		Registry:      sm[12],
		Address:       sm[13],
		Number:        sm[14],
		Neighborhood:  sm[15],
		Reference:     sm[16],
	}
	return result
}
