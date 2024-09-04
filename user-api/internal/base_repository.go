package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// BaseRepository provides common functionality for all repositories.
type BaseRepository struct {
	db *mongo.Database
}

// NewBaseRepository creates a new BaseRepository instance.
func NewBaseRepository(db *mongo.Database) *BaseRepository {
	return &BaseRepository{db: db}
}

// Collection returns a MongoDB collection for the given name.
func (r *BaseRepository) Collection(name string) *mongo.Collection {
	return r.db.Collection(name)
}
