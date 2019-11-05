package database

import (
	"database/sql"
	"fmt"
	"os"
)

var host = os.Getenv("PG_HOST")
var port = os.Getenv("PG_PORT")
var user = os.Getenv("PG_USER")
var password = os.Getenv("PG_PASSWORD")
var dbname = os.Getenv("PG_DB_NAME")
var sslmode = os.Getenv("PG_SSL_MODE")

func Init() *sql.DB {
	psqlinfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}
