package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func A() *fiber.App {
	engine := html.New("./web/pages/A", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/one", func(c *fiber.Ctx) error {
		return c.Render("one/index", nil)
	})

	return app
}
