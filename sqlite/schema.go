package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
  CREATE TABLE IF NOT EXISTS user (
    id   INTEGER     PRIMARY KEY UNIQUE,
    key  varchar(12)  UNIQUE,
    name varchar(20) DEFAULT "",
    level_permision INTEGER DEFAULT 1 
  );
`

type User struct {
	Id             int64  `db:"id"`
	Key            string `db:"key"`
	Name           string `db:"name"`
	LevelPermision int64  `db:"level_permision"`
}
