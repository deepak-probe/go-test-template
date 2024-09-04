package router

import (
	"user-api/internal/user/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app fiber.Router, h *handler.UserHandler) {
	app.Post("/", h.CreateUser)
}
