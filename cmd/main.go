package main

import (
	"github.com/ZombieKiller116/rest-go/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Failed to listen port. \n", err)
	}
}
