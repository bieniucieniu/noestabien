package router

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Router() *fiber.App {
	app := fiber.New()

	app.Mount("/profile", profile())

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Static("/", "./web/templates", fiber.Static{
		Index: "index.html",
	})

	app.Static("/assets", "./web/static")

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/templates/404.html")
	})
	return app
}
