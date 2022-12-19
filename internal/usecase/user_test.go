package usecase_test

import (
	"context"
	"testing"

	"github.com/svbnbyrk/wallet/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/internal/usecase/mocks"
)

func TestUser_Store(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	mockUser := entity.User{
		Name:  "John",
		Email: "john@example.com",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Store", mock.Anything, mockUser).Return(nil)

		uuc := usecase.NewUserUseCase(mockUserRepo)

		err := uuc.Store(context.Background(), mockUser)
		if err != nil {
			println(err)
		}

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}
