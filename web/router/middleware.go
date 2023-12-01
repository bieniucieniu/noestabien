package router

import (
	"github.com/bieniucieniu/noestabien/sqlite"
	"github.com/gofiber/fiber/v2"
)

func levelPermisionLevel(db *sqlite.Database, level int64) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token", "")
		user, err := db.GetUserWithToken(&tokenString)
		if err != nil {
			return err
		}

		if user.LevelPermision < level {
			return c.SendString("not permitted")
		}

		return c.Next()
	}
}
