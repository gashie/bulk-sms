package routes

import (
	"bulk-sms/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/health", controllers.Health)

	//group routes
	GroupRoute(app)

	//contact routes
	ContactRoute(app)

	 //senderid routes
	 SenderRoute(app)

	// //invoice data routes
	// InvoiceDataRoute(app)

	// //invoice detail routes
	// InvoiceDetailsRoute(app)
}
