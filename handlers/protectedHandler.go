package handlers

import "github.com/gofiber/fiber/v2"

func ProtectedRoute(c *fiber.Ctx) error{
	return c.Status(200).JSON(fiber.Map{
		"message" : "this is protected route",
	})
}