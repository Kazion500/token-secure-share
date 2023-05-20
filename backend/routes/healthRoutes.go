package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazion500/secure-share/types"
)

func HealthRoutes(app *fiber.App) {

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(types.ResponseType[string]{
			Success: true,
			Data:    "All GOOD!",
		})
	})
}
