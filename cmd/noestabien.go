package main

import (
	"github.com/bieniucieniu/noestabien/web/router"
)

func main() {
	app := router.Router()

	app.Listen(":3000")
}
