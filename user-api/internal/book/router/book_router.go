package router

import (
	"user-api/internal/book/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(app fiber.Router, h *handler.BookHandler) {
	app.Post("/", h.CreateBook)
}
