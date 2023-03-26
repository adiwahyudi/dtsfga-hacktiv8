package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "adi"
	dbname   = "chap2-challenge3"
)

var (
	DB  *sql.DB
	err error
)

func DatabaseConnection() *sql.DB {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	return DB
}
