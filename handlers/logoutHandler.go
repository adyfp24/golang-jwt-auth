package handlers

import "github.com/gofiber/fiber/v2"

func Logout(c *fiber.Ctx) error{
	c.ClearCookie("token")
	return c.Status(200).JSON(fiber.Map{
		"message":"logout sukses",
	})
}