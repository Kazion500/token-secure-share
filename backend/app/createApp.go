package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazion500/secure-share/routes"
)

func CreateApp() *fiber.App {

	app := fiber.New()

	routes.HealthRoutes(app)
	routes.LinkRoutes(app)

	return app
}
