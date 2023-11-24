package main

import (
	"flag"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	pagesPath := flag.String("pages-path", "./static/pages", "path to static html pages directory")

	app := fiber.New(fiber.Config{UnescapePath: true})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.Dir(*pagesPath),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
	}))

	app.Listen(":3000")
}
