package profile

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

func Router(db *sqlite.Database, baseUrl ...string) *fiber.App {
	engine := html.New("./web/templates/profile", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token", "")
		user, err := db.GetUserWithToken(&tokenString)
		if err != nil {
			log.Println(err)
			return c.Render("profile", fiber.Map{
				"baseUrl": strings.Join(baseUrl, ""),
				"login":   false,
				"id":      "sdasd",
			})
		}
		return c.Render("profile", fiber.Map{
			"baseUrl": strings.Join(baseUrl, ""),
			"login":   true,
			"id":      user.Id,
			"name":    user.Name,
		})
	})

	type Body struct {
		New string `json:"new" xml:"new" form:"new"`
		Key string `json:"key" xml:"key" form:"key"`
	}

	app.Post("/reqUser", func(c *fiber.Ctx) error {
		body := new(Body)
		err := c.BodyParser(body)
		if err != nil {
			log.Println(err)
			return c.SendString(err.Error())
		}

		user := &sqlite.User{}

		if body.Key != "" {
			user, err = db.GetUser(&sqlite.User{
				Key: body.Key,
			})
			if err != nil {
				log.Printf("error getting user by key %s", err.Error())
			}
		}

		if body.New != "" {
			user, err = db.CreateUser(body.New)
			if err != nil {
				log.Printf("error adding user %s", err.Error())
			}
		}

		if user.Id == 0 || user.Key == "" || user.Name == "" {
			return fmt.Errorf("error adding / selecting user \n nil return or invalid user data \n %s", err.Error())
		}

		token, err := auth.CreateToken(&jwt.MapClaims{
			"id":   user.Id,
			"name": user.Name,
		})
		if err != nil {
			return c.Render("reqUser", fiber.Map{
				"name":  user.Name,
				"key":   user.Key,
				"id":    user.Id,
				"error": err.Error(),
			})
		}

		c.Append("Set-Cookie", fmt.Sprintf(`token="%s"`, token))
		return c.Render("reqUser", fiber.Map{
			"name":  user.Name,
			"key":   user.Key,
			"id":    user.Id,
			"token": token,
		})
	})

	return app
}
