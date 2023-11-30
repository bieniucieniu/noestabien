package main

import (
	"log"

	"github.com/bieniucieniu/noestabien/sqlite"
	"github.com/bieniucieniu/noestabien/web/router"
)

func main() {
	db, err := sqlite.New("database.db")
	if err != nil {
		log.Panic(err)
	}

	app := router.Router(db)

	app.Listen(":3000")
}
