package repository

import (
	"context"
	model "user-api/internal/book"

	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book *model.Book) error
}

type bookRepository struct {
	collection *mongo.Collection
}

func NewBookRepository(db *mongo.Database) BookRepository {
	return &bookRepository{collection: db.Collection("books")}
}

func (r *bookRepository) CreateBook(ctx context.Context, book *model.Book) error {
	_, err := r.collection.InsertOne(ctx, book)
	return err
}
