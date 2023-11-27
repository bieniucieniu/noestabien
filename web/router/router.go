package router

import "github.com/gofiber/fiber/v2"

func Router() *fiber.App {
	app := fiber.New()

	app.Static("/assets", "./web/static")

	app.Static("/", "./web/templates", fiber.Static{
		Index: "index.html",
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/templates/404.html")
	})
	return app
}
