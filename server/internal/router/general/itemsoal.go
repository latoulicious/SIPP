package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupItemSoalRoutes(router fiber.Router) {
	itemsoal := router.Group("/item")

	itemsoal.Get("/", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /item")
	})

	itemsoal.Get("/:id", func(c *fiber.Ctx) error {
		// Your logic here
		return c.SendString("GET /item/:id")
	})
}
