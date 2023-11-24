package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	app := fiber.New(fiber.Config{UnescapePath: true})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.Dir("./web/static"),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
	}))

	app.Listen(":3000")
}
