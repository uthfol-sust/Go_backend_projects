package routes

import (
	"urlshorten/database"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ResolverURL(c *fiber.Ctx) error {
	short := c.Params("short") // FIXED

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, short).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in the database",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	return c.Redirect(value, 301)
}
