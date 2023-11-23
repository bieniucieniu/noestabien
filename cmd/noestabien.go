package main

import (
	"github.com/bieniucieniu/noestabien/web/pages"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	index := pages.Index()
	app.Get("/", func(c *fiber.Ctx) error {
		return index.Render(c.Context(), c.Response().BodyWriter())
	})

	app.Listen(":3000")
}
