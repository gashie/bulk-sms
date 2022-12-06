package routes

import (
	"bulk-sms/controllers/settings"

	"github.com/gofiber/fiber/v2"
)

func SenderRoute(app *fiber.App) {

	//status routes
	app.Post("/api/v1/sender", settings.CreateSender)
	app.Get("/api/v1/sender/:id", settings.GetSender)
	app.Get("/api/v1/sender", settings.AllSenders)
	app.Put("/api/v1/sender/:id", settings.UpdateSender)
	app.Delete("/api/v1/sender/:id", settings.DeleteSender)
}
