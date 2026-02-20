package main

import (
	// "fitness_club/database"
	"fitness_club/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// database.Connect()

	app := fiber.New()
	app.Get("/clients", handlers.GetClients)

	app.Get("/clients/:id", handlers.GetClientByID)

	app.Post("/clients", handlers.CreateClient)

	app.Delete("/clients", handlers.DeleteClient)

	app.Put("/clients", handlers.UpdateClient2)

	app.Patch("/clients/:id", handlers.UpdateClient)

	app.Listen(":3000")
}
