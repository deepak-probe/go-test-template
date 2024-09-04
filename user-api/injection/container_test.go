package injection

import (
	"testing"
	"user-api/internal/user/handler"
	"user-api/internal/user/repository"
	"user-api/internal/user/service"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestBuildContainer(t *testing.T) {
	container := BuildContainer()

	t.Run("resolve user handler", func(t *testing.T) {
		err := container.Invoke(func(h *handler.UserHandler) {
			assert.NotNil(t, h)
		})
		assert.NoError(t, err)
	})

	t.Run("resolve user service", func(t *testing.T) {
		err := container.Invoke(func(s service.UserService) {
			assert.NotNil(t, s)
		})
		assert.NoError(t, err)
	})

	t.Run("resolve user repository", func(t *testing.T) {
		err := container.Invoke(func(r repository.UserRepository) {
			assert.NotNil(t, r)
		})
		assert.NoError(t, err)
	})

	t.Run("resolve mongo database", func(t *testing.T) {
		err := container.Invoke(func(db *mongo.Database) {
			assert.NotNil(t, db)
		})
		assert.NoError(t, err)
	})
}
