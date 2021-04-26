package db

import (
	"database/sql"
	"github.com/canmor/go_ms_clean_arch/pkg/util"
)

func Migrate(db *sql.DB) error {
	sqlStmt := `
	create table if not exists blogs (id integer not null primary key, title text, body text, created_at datatime);
	delete from blogs;
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		util.Log().Errorf("%q: %s\n", err, sqlStmt)
	}
	return err
}
