package main

import (
	"flag"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	staticPath := flag.String("static-path", "./web/static", "path to static htmlpages files")

	app := fiber.New(fiber.Config{UnescapePath: true})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.Dir(*staticPath),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
	}))

	app.Listen(":3000")
}
