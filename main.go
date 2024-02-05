package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/adyfp24/golang-jwt-auth/routes"
)

func main(){
	app := fiber.New()
	app.Use(cors.New())
	routes.RouteInit(app)
	err := app.Listen(":4000")
	if err != nil{
		log.Fatal(err)
	}
}