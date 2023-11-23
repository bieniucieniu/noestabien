package web

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func TemplHandler(component templ.Component, options ...func(*templ.ComponentHandler)) fiber.Handler {
	return adaptor.HTTPHandler(templ.Handler(component, options...))
}
