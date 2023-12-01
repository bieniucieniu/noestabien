package router

import (
	"github.com/bieniucieniu/noestabien/sqlite"
	"github.com/gofiber/fiber/v2"
)

func Router(db *sqlite.Database) *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./web/templates/index.html")
	})

	app.Mount("/profile", profile(db, "/profile"))

	app.Static("/assets", "./web/static")

	app.Static("/A", "./web/templates/A", fiber.Static{
		Index: "index.html",
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/templates/404.html")
	})
	return app
}
