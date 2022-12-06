package routes

import (
	"bulk-sms/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/health", controllers.Health)

	//app routes
	 GroupRoute(app)

	// //category routes
	// CategoryRoute(app)

	// //appfields routes
	// AppFieldsRoute(app)

	// //invoice data routes
	// InvoiceDataRoute(app)

	// //invoice detail routes
	// InvoiceDetailsRoute(app)
}
