package main

import (
	"log"

	"github.com/bieniucieniu/noestabien/sqlite"
	"github.com/bieniucieniu/noestabien/web/router"
)

func main() {
	app := router.Router()

	_, err := sqlite.New()
	if err != nil {
		log.Panic(err)
	}

	app.Listen(":3000")
}
