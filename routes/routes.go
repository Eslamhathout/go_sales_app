package routes

import (
	"sales_api_go/go_sales_app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/cashiers/:id/login", controllers.Login)
	// app.Get("/cashiers/:cashierId/logout", controllers.Logout)
	// app.Post("/cashiers/:cashierId/passcode", controllers.Passcode)

	app.Post("/cashier", controllers.CreateCashier)
	app.Get("/cashiers", controllers.CashierList)
	app.Get("/cashier/:id", controllers.GetCashierDetails)
	// app.Put("/cashier/:id", controllers.EditCashier)
	app.Patch("/cashier/:id", controllers.UpdateCashier)
	app.Delete("/cashier/:id", controllers.DeleteCashier)
}
