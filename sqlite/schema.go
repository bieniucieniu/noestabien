package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
  CREATE TABLE IF NOT EXISTS user (
    id   INTEGER     PRIMARY KEY UNIQUE,
    key  varchar(8)  DEFAULT UNIQUE,
    name varchar(20) DEFAULT ""
  );
`

type User struct {
	Id   int     `db:"id"`
	Key  *string `db:"key"`
	Name *string `db:"name"`
}
