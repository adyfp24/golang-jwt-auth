package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ProtectedRoute(c *fiber.Ctx) error{
	userClaims := c.Locals("user")
	if userClaims == nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized: User claims not found",
		})
	}

	token, ok := userClaims.(*jwt.Token)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error: Failed to parse user claims",
		})
	}

	// Melakukan casting klaim-klaim pengguna ke dalam tipe yang sesuai (jwt.MapClaims)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error: Failed to parse user claims",
		})
	}

	// Mengakses informasi pengguna dari klaim-klaim
	userID := claims["user_id"].(float64)
	// Misalnya, untuk mengambil nilai user_id dari klaim-klaim

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Protected route accessed successfully",
		"user_id": userID,
	})
}