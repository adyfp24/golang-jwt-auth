package routes

import (
	"github.com/adyfp24/golang-jwt-auth/handlers"
	"github.com/adyfp24/golang-jwt-auth/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App){
	r := app.Group("/api")
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome in golang jwt auth project ady!")
	})

	// auth route
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)
	r.Get("/protected", middlewares.AuthMiddleware() ,handlers.ProtectedRoute)
	r.Post("/logout", middlewares.AuthMiddleware(), handlers.Logout)

	// protected route

}