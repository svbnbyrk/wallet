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

func TestTransaction_History(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockWalletRepo := new(mocks.WalletRepository)
	mockExchangeRepo := new(mocks.ExchangeRepository)

	mockTransaction := entity.Transaction{
		WalletId:        1,
		TransactionType: "deposit",
		Currency:        "TRY",
		Amount:          1,
	}

	mockListTransaction := make([]entity.Transaction, 0)
	mockListTransaction = append(mockListTransaction, mockTransaction)

	t.Run("success", func(t *testing.T) {
		mockTransactionRepo.On("GetHistory", mock.Anything).Return(mockListTransaction, nil)

		tuc := usecase.NewTransactionUseCase(mockTransactionRepo, mockWalletRepo, mockExchangeRepo)

		ts, err := tuc.History(context.Background())
		if err != nil {
			println(err)
		}

		assert.NoError(t, err)
		assert.Len(t, ts, len(mockListTransaction))

		mockTransactionRepo.AssertExpectations(t)
		mockWalletRepo.AssertExpectations(t)
		mockExchangeRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTransactionRepo.On("GetHistory", mock.Anything).Return(nil, nil)

		tuc := usecase.NewTransactionUseCase(mockTransactionRepo, mockWalletRepo, mockExchangeRepo)

		ts, err := tuc.History(context.Background())
		if err != nil {
			println(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, ts, nil)

		mockTransactionRepo.AssertExpectations(t)
		mockWalletRepo.AssertExpectations(t)
		mockExchangeRepo.AssertExpectations(t)
	})
}
