package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/svbnbyrk/wallet/internal/core/usecase"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/internal/entity/mocks"
)

func TestUser_Store(t *testing.T) {
	mockUser := entity.User{
		Name:  "John",
		Email: "john@example.com",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("Store", mock.Anything, mockUser).Return(nil)

		uuc := usecase.NewUserUseCase(mockUserRepo)

		err := uuc.Store(context.Background(), mockUser)

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("Store", mock.Anything, mockUser).Return(entity.ErrInternalServerError)

		uuc := usecase.NewUserUseCase(mockUserRepo)

		err := uuc.Store(context.Background(), mockUser)

		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}
