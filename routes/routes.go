package routes

import(
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App){
	r := app.Group("/api")
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome in golang jwt auth project ady!")
	})
}