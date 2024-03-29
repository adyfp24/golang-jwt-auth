package handlers

import (
	"github.com/adyfp24/golang-jwt-auth/database"
	"github.com/adyfp24/golang-jwt-auth/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error{
	db := database.DB
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "bad gateway",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to hash password",
			"error":   err.Error(),
		})
	}

	user.Password = string(hashedPassword)

	result := db.Create(&user).Error
	if result != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "gagal menambahkan data user",
			"error": err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "berhasil menambahkan data user",
		"data": user,
	})
}