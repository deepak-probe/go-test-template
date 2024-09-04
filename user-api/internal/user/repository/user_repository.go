package repository

import (
	"context"
	base_repository "user-api/internal"
	model "user-api/internal/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository interface defines methods for user repository.
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByName(ctx context.Context, name string) (*model.User, error) // New method
}

type userRepository struct {
	base_repository.BaseRepository
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{BaseRepository: *base_repository.NewBaseRepository(db)}
}

// CreateUser inserts a new user into the database.
func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.Collection("users").InsertOne(ctx, user)
	return err
}

// GetUserByName fetches a user by their name.
func (r *userRepository) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	err := r.Collection("users").FindOne(ctx, bson.M{"name": name}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No user found
		}
		return nil, err
	}
	return &user, nil
}
