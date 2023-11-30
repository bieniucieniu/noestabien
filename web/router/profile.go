package router

import (
	"fmt"
	"log"
	"strings"

	"github.com/bieniucieniu/noestabien/auth"
	"github.com/bieniucieniu/noestabien/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/golang-jwt/jwt/v5"
)

func profile(db *sqlite.Database, baseUrl ...string) *fiber.App {
	engine := html.New("./web/templates/profile", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token", "")
		user, err := db.GetUserWithToken(&tokenString)
		if err != nil {
			log.Println(err)
			return c.Render("index", fiber.Map{
				"baseUrl": strings.Join(baseUrl, ""),
				"login":   false,
				"id":      "sdasd",
			})
		}
		return c.Render("index", fiber.Map{
			"baseUrl": strings.Join(baseUrl, ""),
			"login":   true,
			"id":      user.Id,
			"name":    user.Name,
		})
	})

	type Body struct {
		Name string `json:"name" xml:"name" form:"name"`
	}

	app.Post("/genUser", func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token", "")
		_, err := auth.ValidateToken(&tokenString)
		if err != nil {
			return c.SendString("valid token already present")
		}

		body := new(Body)
		if err := c.BodyParser(body); err != nil {
			log.Println(err)
			return c.SendString(err.Error())
		}

		user, err := db.CreateUser(body.Name)
		if err != nil {
			c.SendString(fmt.Sprintf("error adding user %s", err.Error()))
		}

		token, err := auth.CreateToken(&jwt.MapClaims{
			"id":   user.Id,
			"name": user.Name,
		})
		if err != nil {
			return c.Render("genUser", fiber.Map{
				"name":  user.Name,
				"key":   user.Key,
				"id":    user.Id,
				"error": err.Error(),
			})
		}

		c.Append("Set-Cookie", fmt.Sprintf(`token="%s"`, token))
		return c.Render("genUser", fiber.Map{
			"name":  user.Name,
			"key":   user.Key,
			"id":    user.Id,
			"token": token,
		})
	})

	return app
}
