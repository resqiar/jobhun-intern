package main

import (
	"jobhun-intern/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init server
	server := fiber.New()

	// init routes
	routes.RouteInit(server)

	// Listen server on PORT ...
	PORT := os.Getenv("PORT")
	server.Listen(":" + PORT)
}
