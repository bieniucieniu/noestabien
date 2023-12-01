package router

import (
	"github.com/bieniucieniu/noestabien/sqlite"
	"github.com/bieniucieniu/noestabien/web/router/profile"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router(db *sqlite.Database) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Set-Cookie",
		MaxAge:           1000,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./web/templates/index.html")
	})

	app.Static("/assets", "./web/static")

	app.Mount("/profile", profile.Router(db, "/profile"))

	app.Use("/A/*", levelPermisionLevel(db, 1))
	app.Static("/A", "./web/templates/A", fiber.Static{
		Index: "index.html",
	})

	app.Get("/**/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/templates/404.html")
	})
	return app
}
