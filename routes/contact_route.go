package routes

import (
	"bulk-sms/controllers/settings"

	"github.com/gofiber/fiber/v2"
)

func ContactRoute(app *fiber.App) {

	//status routes
	app.Post("/api/v1/contact", settings.CreateContact)
	app.Get("/api/v1/contact/:id", settings.GetContact)
	app.Get("/api/v1/contact", settings.AllContacts)
	app.Put("/api/v1/contact/:id", settings.UpdateContact)
	app.Delete("/api/v1/contact/:id", settings.DeleteContact)
}
