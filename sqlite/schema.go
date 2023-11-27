package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
  CREATE TABLE IF NOT EXISTS user (
    id   INTEGER     PRIMARY KEY,
    key  varchar(8)  DEFAULT "",
    name varchar(20) DEFAULT ""
  );
`

type User struct {
	Id  string `db:"id"`
	Key string `db:"key"`
}
