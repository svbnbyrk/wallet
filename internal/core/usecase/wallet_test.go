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

func TestWallet_Store(t *testing.T) {
	mockWallet := entity.Wallet{
		UserId:   1,
		Balance:  1,
		Currency: "TRY",
	}

	t.Run("success", func(t *testing.T) {
		mockWalletRepo := new(mocks.WalletRepository)
		mockWalletRepo.On("Store", mock.Anything, mockWallet).Return(nil)

		wuc := usecase.NewWalletUseCase(mockWalletRepo)

		err := wuc.Store(context.Background(), mockWallet)

		assert.NoError(t, err)

		mockWalletRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockWalletRepo := new(mocks.WalletRepository)
		mockWalletRepo.On("Store", mock.Anything, mockWallet).Return(entity.ErrInternalServerError)

		wuc := usecase.NewWalletUseCase(mockWalletRepo)

		err := wuc.Store(context.Background(), mockWallet)

		assert.Error(t, err)

		mockWalletRepo.AssertExpectations(t)
	})
}

func TestWallet_GetWalletsbyUser(t *testing.T) {
	mockWallet := entity.Wallet{
		UserId:   1,
		Balance:  1,
		Currency: "TRY",
	}

	mockListWallet := make([]entity.Wallet, 0)
	mockListWallet = append(mockListWallet, mockWallet)

	t.Run("success", func(t *testing.T) {
		mockWalletRepo := new(mocks.WalletRepository)

		mockWalletRepo.On("GetWalletsByUser", mock.Anything, mock.AnythingOfType("int64")).Return(mockListWallet, nil).Once()

		wuc := usecase.NewWalletUseCase(mockWalletRepo)

		num := int64(1)
		ws, err := wuc.GetWalletsByUser(context.Background(), num)

		assert.NoError(t, err)
		assert.Len(t, ws, len(mockListWallet))

		mockWalletRepo.AssertExpectations(t)
	})
}
