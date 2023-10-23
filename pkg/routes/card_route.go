package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thabish/go-rest/pkg/controllers"
)

func CardRoute(app *fiber.App) {
	app.Post("/card", controllers.CreateCard)
	app.Get("/card/:id", controllers.GetCard)
	app.Get("/cards", controllers.GetAllCards)
	app.Delete("/card/:id", controllers.DeleteCard)
	app.Put("/card/:id", controllers.UpdateCard)
}