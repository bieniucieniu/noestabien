package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func A() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/A/index.html")
	})

	app.Get("/one", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/A/one/index.html")
	})

	app.Get("/two", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/A/two/index.html")
	})

	app.Get("/here", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/A/here/index.html")
	})

	app.Get("/quark", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/A/quark/index.html")
	})

	app.Get("/move", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/A/move/index.html")
	})

	app.Get("/checkmate", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/A/checkmate/index.html")
	})

	app.Get("/b/checkmate", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/A/b/checkmate/index.html")
	})

	app.Get("/qxf7", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/A/qxf7/index.html")
	})

	return app
}
