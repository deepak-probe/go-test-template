package router

import (
	"user-api/internal/user/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	app.Post("/users", userHandler.CreateUser)
}
