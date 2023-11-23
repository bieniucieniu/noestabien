package main

import (
	"github.com/bieniucieniu/noestabien/web"
	"github.com/bieniucieniu/noestabien/web/pages/templ"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	index := templ.Index()
	app.Get("/", web.TemplHandler(index))

	app.Listen(":3000")
}
