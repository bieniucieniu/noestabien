package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func profile() *fiber.App {
	engine := html.New("./web/templates/profile", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"login": true,
			"id":    "sdasd",
		})
	})

	return app
}
