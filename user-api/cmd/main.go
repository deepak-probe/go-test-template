package main

import (
	"log"
	"user-api/injection"
	book_handler "user-api/internal/book/handler"
	book_router "user-api/internal/book/router"
	user_handler "user-api/internal/user/handler"
	user_router "user-api/internal/user/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Build the DI container
	container := injection.BuildContainer()

	// Invoke dependencies and start the server
	if err := container.Invoke(func(userHandler *user_handler.UserHandler, bookHandler *book_handler.BookHandler) error {
		app := fiber.New()

		// Create a common prefix group
		api := app.Group("/api")

		// Setup user routes under /api/users
		userGroup := api.Group("/users")
		user_router.SetupUserRoutes(userGroup, userHandler)

		// Setup book routes under /api/books
		bookGroup := api.Group("/books")
		book_router.SetupBookRoutes(bookGroup, bookHandler)

		// Start the server
		if err := app.Listen(":8080"); err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalf("Failed to invoke dependencies or start server: %v", err)
	}
}
