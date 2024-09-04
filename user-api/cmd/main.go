package main

import (
	"log"
	"user-api/internal/di"
	"user-api/internal/user/handler"
	"user-api/internal/user/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(userHandler *handler.UserHandler) {
		app := fiber.New()

		// Setup routes
		router.SetupRoutes(app, userHandler)

		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Failed to invoke dependencies: %v", err)
	}
}
