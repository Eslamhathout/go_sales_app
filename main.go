package main

import (
	db "sales_api_go/go_sales_app/config"
	routes "sales_api_go/go_sales_app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()

	routes.Setup(app)
	app.Listen(":3000")

}
