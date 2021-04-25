package db

import (
	_ "github.com/mattn/go-sqlite3"
	"log"

	"database/sql"
)

func NewInMemory() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
