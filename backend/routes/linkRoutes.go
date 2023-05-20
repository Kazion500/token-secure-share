package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazion500/secure-share/controllers"
)

func LinkRoutes(app *fiber.App) {
	link := app.Group("links")

	link.Post("/", controllers.CreateLink)
	link.Get("/:linkId", controllers.GetLink)
}
