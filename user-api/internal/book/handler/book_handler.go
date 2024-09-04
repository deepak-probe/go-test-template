package handler

import (
	"context"
	model "user-api/internal/book"
	"user-api/internal/book/service"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(svc service.BookService) *BookHandler {
	return &BookHandler{bookService: svc}
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var book model.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.bookService.CreateBook(context.Background(), &book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}
