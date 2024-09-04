package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"user-api/internal/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockUserService(ctrl)
	userHandler := NewUserHandler(mockService)

	app := fiber.New()
	app.Post("/users", userHandler.CreateUser)

	userJSON := `{"name": "John Doe", "email": "johndoe@example.com", "password": "securepassword"}`

	t.Run("successful user creation", func(t *testing.T) {
		mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("user creation failure", func(t *testing.T) {
		mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(assert.AnError)

		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}
