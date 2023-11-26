package main

import (
	"github.com/bieniucieniu/noestabien/web/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/index.html")
	})

	app.Mount("/A/", handlers.A())

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/404.html")
	})

	app.Listen(":3000")
}
