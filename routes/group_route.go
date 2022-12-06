package routes

import (
	"bulk-sms/controllers/settings"

	"github.com/gofiber/fiber/v2"
)

func GroupRoute(app *fiber.App) {
	
    //status routes
	app.Post("/api/v1/group", settings.CreateGroup)
	app.Get("/api/v1/group/:id", settings.GetGroup)
	app.Get("/api/v1/group", settings.AllGroups)
	app.Put("/api/v1/group/:id", settings.UpdateGroup)
	app.Delete("/api/v1/group/:id", settings.DeleteGroup)
}
