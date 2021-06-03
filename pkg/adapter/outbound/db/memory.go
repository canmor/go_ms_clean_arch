package db

import (
	"database/sql"
	"github.com/canmor/go_ms_clean_arch/pkg/util"
	_ "github.com/mattn/go-sqlite3"
)

func NewInMemory() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		util.Log().Error(err)
	}
	return db, err
}
