package service

import (
	"context"
	"fmt"
	model "user-api/internal/book"
	"user-api/internal/book/repository"
	user_repo "user-api/internal/user/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, book *model.Book) error
	AuthorExists(ctx context.Context, authorName string) (bool, error) // Updated method
}

type bookService struct {
	repo     repository.BookRepository
	userRepo user_repo.UserRepository // Add user repository
}

func NewBookService(repo repository.BookRepository, userRepo user_repo.UserRepository) BookService {
	return &bookService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *bookService) CreateBook(ctx context.Context, book *model.Book) error {
	exists, err := s.AuthorExists(ctx, book.Author)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("author %s does not exist", book.Author)
	}
	return s.repo.CreateBook(ctx, book)
}

// AuthorExists checks if an author exists in the user repository.
func (s *bookService) AuthorExists(ctx context.Context, authorName string) (bool, error) {
	fmt.Println("this is me1: ", authorName)
	user, err := s.userRepo.GetUserByName(ctx, authorName)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
