package injection

import (
	"context"
	book_handler "user-api/internal/book/handler"
	book_repository "user-api/internal/book/repository"
	book_service "user-api/internal/book/service"

	user_handler "user-api/internal/user/handler"
	user_repository "user-api/internal/user/repository"
	user_service "user-api/internal/user/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Provide the MongoDB database instance
	container.Provide(func() (*mongo.Database, error) {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://deepak:deepak123@localhost:8001"))
		if err != nil {
			return nil, err
		}
		return client.Database("userdb"), nil
	})

	// Provide User Repository
	container.Provide(user_repository.NewUserRepository)
	// Provide User Service
	container.Provide(user_service.NewUserService)
	// Provide User Handler
	container.Provide(user_handler.NewUserHandler)

	// Provide Book Repository
	container.Provide(book_repository.NewBookRepository)
	// Provide Book Service
	container.Provide(func(bookRepo book_repository.BookRepository, userRepo user_repository.UserRepository) book_service.BookService {
		return book_service.NewBookService(bookRepo, userRepo)
	})
	// Provide Book Handler
	container.Provide(book_handler.NewBookHandler)

	return container
}
