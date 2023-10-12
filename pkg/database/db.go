package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDb() error {
	d, err := sqlx.Connect("postgres", "user=dev password=abcde dbname=stocktrading sslmode=disable")
	if err != nil {
		return err
	}
	db = d
	db.MustExec(schema)
	return nil
}
