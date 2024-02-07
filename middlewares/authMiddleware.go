package middlewares

import (
	"strings"
	"github.com/adyfp24/golang-jwt-auth/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler{
	return func(c *fiber.Ctx) error{

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized: No token provided",
			})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized: Invalid authorization header format",
			})
		}
		token := tokenParts[1]

		if token == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		c.Locals("user", claims)

		return c.Next()
	}
}