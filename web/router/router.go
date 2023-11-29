package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Router(db *sqlx.DB) *fiber.App {
	app := fiber.New()

	app.Mount("/profile", profile(db, "/profile"))

	app.Static("/", "./web/templates", fiber.Static{
		Index: "index.html",
	})

	app.Static("/assets", "./web/static")

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/templates/404.html")
	})
	return app
}
