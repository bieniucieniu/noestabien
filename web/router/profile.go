package router

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bieniucieniu/noestabien/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jmoiron/sqlx"

	"github.com/golang-jwt/jwt/v5"
)

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func profile(db *sqlx.DB, baseUrl ...string) *fiber.App {
	engine := html.New("./web/templates/profile", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"baseUrl": strings.Join(baseUrl, ""),
			"login":   false,
			"id":      "sdasd",
		})
	})

	app.Post("/genUser/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name", "")
		key := fmt.Sprintf("%08d", rand.Intn(100000000))

		db.MustExec("INSERT INTO user (id, key, name) values ($1, $2, $3)", nil, key, name)

		user := sqlite.User{}
		db.Get(&user, "SELECT * FROM user WHERE key=$1", key)

		claims := jwt.MapClaims{
			"id":   user.Id,
			"name": user.Name,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("jajcocjaj"))
		if err != nil {
			return c.Render("genUser", fiber.Map{
				"token": "",
				"name":  name,
				"key":   key,
				"error": err.Error(),
			})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    t,
			HTTPOnly: true,
		})

		return c.Render("genUser", fiber.Map{
			"token": t,
			"name":  name,
			"key":   key,
		})
	})

	return app
}
