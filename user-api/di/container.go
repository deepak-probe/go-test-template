package di

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

	container.Provide(func() (*mongo.Database, error) {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			return nil, err
		}
		return client.Database("userdb"), nil
	})

	// User dependencies
	container.Provide(user_repository.NewUserRepository)
	container.Provide(user_service.NewUserService)
	container.Provide(user_handler.NewUserHandler)

	// Book dependencies
	container.Provide(book_repository.NewBookRepository)
	container.Provide(book_service.NewBookService)
	container.Provide(book_handler.NewBookHandler)

	return container
}
