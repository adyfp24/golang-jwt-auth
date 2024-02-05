package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/adyfp24/golang-jwt-auth/routes"
	"github.com/adyfp24/golang-jwt-auth/database"
)

func main(){
	database.InitDB()
	app := fiber.New()
	app.Use(cors.New())
	routes.RouteInit(app)
	err := app.Listen(":4000")
	if err != nil{
		log.Fatal(err)
	}
}