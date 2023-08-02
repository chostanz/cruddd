package database

import (
	sqlx "github.com/jmoiron/sqlx"
)

var DB = Connection()

func Connection() *sqlx.DB {
	db, _ := sqlx.Connect("postgres", "user=postgres password=12345 dbname= db_users sslmode=disable")

	return db
}
