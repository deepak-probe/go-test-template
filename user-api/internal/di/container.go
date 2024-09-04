package di

import (
	"context"
	"user-api/internal/user/handler"
	"user-api/internal/user/repository"
	"user-api/internal/user/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(func() (*mongo.Database, error) {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://deepak:deepak123@localhost:8001"))
		if err != nil {
			return nil, err
		}
		return client.Database("userdb"), nil
	})

	container.Provide(repository.NewUserRepository)
	container.Provide(service.NewUserService)
	container.Provide(handler.NewUserHandler)

	return container
}
