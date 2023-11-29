package sqlite

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bieniucieniu/noestabien/auth"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	sqlx *sqlx.DB
}

func New() (*Database, error) {
	sql, err := sqlx.Connect("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	sql.MustExec(schema)

	db := Database{
		sqlx: sql,
	}

	return &db, nil
}

// todo
func (db *Database) AddUser(u *User) (*User, error) {
	user := new(User)
	out := db.sqlx.MustExec("INSERT INTO user (id, key, name) values ($1, $2, $3) RETURNING", u.Id, u.Key, u.Name)

	fmt.Printf("out: %v\n", out)
	return user, nil
}

func (db *Database) GetUser(id int64) (*User, error) {
	user := new(User)
	db.sqlx.Get(&user, "SELECT * FROM user WHERE id=$1", id)
	return user, nil
}

func (db *Database) GetUserWithToken(tokenString *string) (*User, error) {
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	idStr, ok := (*claims)["id"]
	if !ok {
		return nil, fmt.Errorf("no id in token")
	}
	id, err := strconv.ParseInt(fmt.Sprint(idStr), 10, 64)
	if err != nil {
		log.Printf("error %s", err.Error())
		return nil, err
	}

	user, err := db.GetUser(id)
	if err != nil {
		log.Printf("error %s", err.Error())
		return nil, err
	}

	return user, nil
}
