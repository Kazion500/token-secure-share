package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kazion500/secure-share/routes"
)

func CreateApp() *fiber.App {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://127.0.0.1:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.HealthRoutes(app)
	routes.LinkRoutes(app)

	return app
}
