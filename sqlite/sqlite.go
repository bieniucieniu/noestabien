package sqlite

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func New() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	db.MustExec(schema)

	return db, nil
}
