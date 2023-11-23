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

	one := templ.One()
	app.Get("/one", web.TemplHandler(one))

	here := templ.Here()
	app.Get("/here", web.TemplHandler(here))

	notFound := templ.NotFound()
	app.Get("/*", web.TemplHandler(notFound))

	app.Listen(":3000")
}
