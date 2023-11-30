package router

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

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
			log.Print(err)
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
		body := new(Body)

		if err := c.BodyParser(body); err != nil {
			log.Println(err)
			return c.SendString(err.Error())
		}

		if l := len(body.Name); l > 20 {
			return c.SendString("to long name")
		} else if l == 0 {
			return c.SendString("no name provided")
		}

		key := fmt.Sprintf("%12d", rand.Intn(1_000_000_000_000))

		user, err := db.AddUser(&sqlite.User{
			Id:   nil,
			Name: body.Name,
			Key:  key,
		})
		if err != nil {
			c.SendString(fmt.Sprintf("error adding user %s", err.Error()))
		}

		claims := jwt.MapClaims{
			"id":   user.Id,
			"name": user.Name,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("jajcocjaj"))
		if err != nil {
			return c.Render("genUser", fiber.Map{
				"token": "",
				"name":  user.Name,
				"key":   user.Key,
				"id":    user.Id,
				"error": err.Error(),
			})
		}

		return c.Render("genUser", fiber.Map{
			"name":  user.Name,
			"key":   user.Key,
			"id":    user.Id,
			"token": t,
		})
	})

	return app
}
