package main

import (
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

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Listen server on PORT ...
	PORT := os.Getenv("PORT")
	server.Listen(":" + PORT)
}
