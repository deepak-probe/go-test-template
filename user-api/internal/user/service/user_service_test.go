package service

import (
	"context"
	"testing"
	model "user-api/internal/user"
	"user-api/internal/user/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userSvc := NewUserService(mockRepo)

	user := &model.User{Name: "John Doe", Email: "john.doe@example.com", Password: "securepassword"}

	t.Run("successful user creation", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(gomock.Any(), user).Return(nil)

		err := userSvc.CreateUser(context.Background(), user)
		assert.NoError(t, err)
	})

	t.Run("user creation failure", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(gomock.Any(), user).Return(assert.AnError)

		err := userSvc.CreateUser(context.Background(), user)
		assert.Error(t, err)
	})
}
