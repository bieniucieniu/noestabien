package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func A() *fiber.App {
	engine := html.New("./web/pages/A", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/one")
	})

	app.Get("/one", func(c *fiber.Ctx) error {
		return c.Render("one/index", nil)
	})

	app.Get("/two", func(c *fiber.Ctx) error {
		return c.Render("two/index", nil)
	})

	app.Get("/here", func(c *fiber.Ctx) error {
		return c.Render("here/index", nil)
	})

	app.Get("/quark", func(c *fiber.Ctx) error {
		return c.Render("quark/index", nil)
	})

	app.Get("/move", func(c *fiber.Ctx) error {
		return c.Render("move/index", nil)
	})

	app.Get("/checkmate", func(c *fiber.Ctx) error {
		return c.Render("checkmate/index", nil)
	})

	app.Get("/b/checkmate", func(c *fiber.Ctx) error {
		return c.Render("b/checkmate/index", nil)
	})

	app.Get("/qxf7", func(c *fiber.Ctx) error {
		return c.Render("qxf7/index", nil)
	})

	return app
}
