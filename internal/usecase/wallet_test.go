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

func TestWallet_Store(t *testing.T) {
	mockWalletRepo := new(mocks.WalletRepository)

	mockWallet := entity.Wallet{
		UserId:   1,
		Balance:  1,
		Currency: "TRY",
	}

	t.Run("success", func(t *testing.T) {
		mockWalletRepo.On("Store", mock.Anything, mockWallet).Return(nil)

		wuc := usecase.NewWalletUseCase(mockWalletRepo)

		err := wuc.Store(context.Background(), mockWallet)
		if err != nil {
			println(err)
		}

		assert.NoError(t, err)

		mockWalletRepo.AssertExpectations(t)
	})
}

func TestWallet_GetWalletsbyUser(t *testing.T) {
	mockWalletRepo := new(mocks.WalletRepository)

	mockWallet := entity.Wallet{
		UserId:   1,
		Balance:  1,
		Currency: "TRY",
	}

	mockListWallet := make([]entity.Wallet, 0)
	mockListWallet = append(mockListWallet, mockWallet)

	t.Run("success", func(t *testing.T) {
		mockWalletRepo.On("GetWalletsbyUser", mock.Anything, mock.AnythingOfType("int")).Return(mockListWallet, nil)

		wuc := usecase.NewWalletUseCase(mockWalletRepo)

		num := int(1)
		ws, err := wuc.GetWalletsbyUser(context.Background(), num)

		assert.NoError(t, err)
		assert.Len(t, ws, len(mockListWallet))

		mockWalletRepo.AssertExpectations(t)
	})
}
