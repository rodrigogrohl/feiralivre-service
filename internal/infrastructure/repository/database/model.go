package database

import (
	"context"

	"github.com/rodrigogrohl/feiralivre-service/pkg/canonical"
	"github.com/sirupsen/logrus"
)

func ResetAllModel(dbStruct *DBStruct) error {
	if dbStruct.DB != nil {
		err := dbStruct.DB.ResetModel(context.Background(), (*canonical.StreetMarket)(nil))
		if err != nil {
			return err
		}
	} else {
		model := `
			CREATE TABLE IF NOT EXISTS street_market (
				id INTEGER NOT NULL PRIMARY KEY,
				longitude INTEGER,
				latitude INTEGER,
				sector_cense INTEGER,
				area_ponderate INTEGER,
				district_code INTEGER,
				district TEXT,
				sub_town_code INTEGER,
				sub_town TEXT,
				region5 TEXT,
				region8 TEXT,
				name_alias TEXT,
				registry TEXT,
				addr TEXT,
				addr_number TEXT,
				neighborhood TEXT,
				reference TEXT
			);
		`
		res, err := dbStruct.Conn.Exec(model)
		if err != nil {
			return err
		}
		logrus.WithField("response", res).Info("model was configured")
	}
	return nil
}
