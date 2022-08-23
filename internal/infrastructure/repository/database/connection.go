package database

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rodrigogrohl/feiralivre-service/internal/infrastructure/config"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var (
	_onceConnection sync.Once
	_dbStruct       *DBStruct
	_bunDB          *bun.DB
)

func DatabaseConnect() *DBStruct {
	_onceConnection.Do(func() {
		connString := config.DbConnection
		connDriver := config.DbDriver

		if connString == "" {
			logrus.Panic("received empty database connection string")
		}

		logrus.WithField("conn", connString).WithField("driver", connDriver).Info("connecting")
		conn, err := sql.Open(connDriver, connString)
		CheckError(err)

		if connDriver == "postgres" {
			_bunDB = bun.NewDB(conn, pgdialect.New())
		}

		_dbStruct = &DBStruct{
			Conn: conn,
			DB:   _bunDB,
		}
	})
	return _dbStruct
}

func CheckError(err error) {
	if err != nil {
		logrus.Panic(err)
	}
}
