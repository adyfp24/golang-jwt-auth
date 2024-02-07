package handlers

import (
	"github.com/adyfp24/golang-jwt-auth/database"
	"github.com/adyfp24/golang-jwt-auth/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/adyfp24/golang-jwt-auth/utils"
)

func Login(c *fiber.Ctx) error{
	db := database.DB
	loginRequest := new(models.LoginRequest)
	err := c.BodyParser(loginRequest)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "bad gateway",
			"error":   err.Error(),
		})
	} 
	
	user := new(models.User)
	err = db.Where("username = ?", loginRequest.Username).First(user).Error
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
            "message": "username belum terdaftar",
            "error":   err.Error(),
        })
	}				
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil{
		return c.Status(401).JSON(fiber.Map{
            "message": "password tidak sesuai",
            "error":   err.Error(),
        })
	}

	token, err := utils.GenerateToken(user.ID)

	return c.Status(201).JSON(fiber.Map{
        "message": "Login successful",
        "token": token,
    })
}