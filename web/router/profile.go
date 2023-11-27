package router

import (
	"fmt"
	"math/rand"

	"github.com/bieniucieniu/noestabien/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jmoiron/sqlx"

	"github.com/golang-jwt/jwt/v5"
)

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func profile(db *sqlx.DB) *fiber.App {
	engine := html.New("./web/templates/profile", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"login": true,
			"id":    "sdasd",
		})
	})

	app.Post("/genUser", func(c *fiber.Ctx) error {
		key := fmt.Sprintf("%08d", rand.Intn(100000000))

		db.MustExec("INSERT INTO user (id, key, name) values ($1, $2, $3)", nil, key, fmt.Sprintf("%08d", rand.Intn(100000000)))

		user := sqlite.User{}
		db.Get(&user, "SELECT * FROM user WHERE key=$1", key)

		claims := jwt.MapClaims{
			"id":   user.Id,
			"name": user.Name,
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("jajcocjaj"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendString(key + " " + t)
	})
	return app
}
