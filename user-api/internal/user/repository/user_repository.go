package repository

import (
	"context"
	model "user-api/internal/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.db.Collection("users").InsertOne(ctx, user)
	return err
}
