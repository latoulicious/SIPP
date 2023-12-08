// internal/cors/cors.go
package cors

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupCors configures CORS for the application
func SetupCors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))
}
