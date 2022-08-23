package database

import (
	"database/sql"

	"github.com/uptrace/bun"
)

type DBStruct struct {
	Conn *sql.DB
	DB   *bun.DB
}
