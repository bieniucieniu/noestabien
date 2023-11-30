package sqlite

import (
	"fmt"
	"log"
	"math/rand"

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

func (db *Database) CreateUser(name string) (*User, error) {
	if l := len(name); l > 20 {
		return nil, fmt.Errorf("to long name")
	} else if l == 0 {
		return nil, fmt.Errorf("no name provided")
	}

	key := fmt.Sprintf("%12d", rand.Intn(1_000_000_000_000))

	user, err := db.AddUser(&User{
		Name: name,
		Key:  key,
	})
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("error adding user %s", err.Error()))
	}
	return user, nil
}

func (db *Database) AddUser(u *User) (*User, error) {
	user := new(User)
	db.sqlx.QueryRowx("INSERT INTO user (id, key, name) values ($1, $2, $3) RETURNING *;", u.Id, u.Key, u.Name).StructScan(user)
	return user, nil
}

func (db *Database) GetUser(id int64) (*User, error) {
	user := User{}
	err := db.sqlx.Get(&user, "SELECT * FROM user WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *Database) GetUserWithToken(tokenString *string) (*User, error) {
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	idStr, ok := (*claims)["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("no id in token")
	}

	user, err := db.GetUser(int64(idStr))
	if err != nil {
		log.Printf("error %s", err.Error())
		return nil, err
	}

	return user, nil
}
