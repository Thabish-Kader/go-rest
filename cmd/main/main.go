package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thabish/go-rest/pkg/routes"
)




func main() {
	app := fiber.New()
	// config.ConnectDb()
	routes.CardRoute(app)
	app.Listen(":8000")


}