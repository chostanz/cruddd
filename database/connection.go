package database

import (
	"log"

	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB = Connection()

func Connection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=12345 dbname= db_users sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
