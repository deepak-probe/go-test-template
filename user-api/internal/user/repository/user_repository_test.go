package repository

import (
	"context"
	"testing"
	model "user-api/internal/user"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mt.Finish()

	mt.Run("successful user creation", func(mt *mtest.T) {
		userRepo := NewUserRepository(mt.DB)

		user := &model.User{Name: "John Doe", Email: "john.doe@example.com", Password: "securepassword"}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := userRepo.CreateUser(context.Background(), user)

		assert.NoError(t, err)
		mt.ClearMockResponses()
	})

	mt.Run("database insertion error", func(mt *mtest.T) {
		userRepo := NewUserRepository(mt.DB)

		user := &model.User{Name: "John Doe", Email: "john.doe@example.com", Password: "securepassword"}

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))

		err := userRepo.CreateUser(context.Background(), user)

		assert.Error(t, err)
	})
}
