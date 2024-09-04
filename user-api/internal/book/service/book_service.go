package service

import (
	"context"
	model "user-api/internal/book"
	"user-api/internal/book/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, book *model.Book) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(ctx context.Context, book *model.Book) error {
	return s.repo.CreateBook(ctx, book)
}
