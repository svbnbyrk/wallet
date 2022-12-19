package usecase_test

import (
	"context"
	"testing"

	"github.com/svbnbyrk/wallet/internal/core/usecase"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/internal/entity/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		mockTransactionRepo.On("GetHistory", mock.Anything).Return(mockListTransaction, nil).Once()

		tuc := usecase.NewTransactionUseCase(mockTransactionRepo, mockWalletRepo, mockExchangeRepo)

		history, err := tuc.History(context.Background())

		assert.NoError(t, err)
		assert.Len(t, history, len(mockListTransaction))

		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockTransactionRepo.On("GetHistory", mock.Anything).Return(nil, nil).Once()

		tuc := usecase.NewTransactionUseCase(mockTransactionRepo, mockWalletRepo, mockExchangeRepo)

		history, err := tuc.History(context.Background())

		assert.EqualError(t, err, entity.ErrNotFound.Error())
		assert.Len(t, history, 0)
		assert.Nil(t, history)

		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTransactionRepo.On("GetHistory", mock.Anything).Return(nil, entity.ErrInternalServerError).Once()

		tuc := usecase.NewTransactionUseCase(mockTransactionRepo, mockWalletRepo, mockExchangeRepo)

		history, err := tuc.History(context.Background())

		assert.Len(t, history, 0)
		assert.Error(t, err)
		assert.Nil(t, history)

		mockTransactionRepo.AssertExpectations(t)
	})
}
func TestTransaction_Post(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockWalletRepo := new(mocks.WalletRepository)
	mockExchangeRepo := new(mocks.ExchangeRepository)

	t.Run("success transaction of deposit with same currency", func(t *testing.T) {
		mockTransaction := entity.Transaction{
			WalletId:        1,
			TransactionType: "deposit",
			Currency:        "TRY",
			Amount:          10,
		}

		mockWallet := entity.Wallet{
			Currency: "TRY",
			Balance:  0,
			UserId:   1,
		}

		mockExchange := entity.Exchange{
			Currency: "TRY",
			Rate:     1,
		}

		mockWalletRepo.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(mockWallet, nil).Once()

		mockExchangeRepo.On("GetByCurrency", mock.Anything, "TRY").Return(mockExchange, nil)

		//if is balance correct value return true
		mockWalletRepo.On("Update", mock.Anything, mock.MatchedBy(func(w entity.Wallet) bool {
			return w.Balance == -10
		})).Return(nil)

		mockTransactionRepo.On("Store", mock.Anything, mock.MatchedBy(func(t entity.Transaction) bool {
			return t.TransactionType == "deposit"
		})).Return(nil)

		tuc := usecase.NewTransactionUseCase(mockTransactionRepo, mockWalletRepo, mockExchangeRepo)

		err := tuc.Post(context.Background(), mockTransaction)
		if err != nil {
			println(err)
		}

		assert.NoError(t, err)

		mockWalletRepo.AssertExpectations(t)
		mockExchangeRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("success transaction of withdraw with same currency", func(t *testing.T) {
		mockTransaction := entity.Transaction{
			WalletId:        1,
			TransactionType: "withdraw",
			Currency:        "TRY",
			Amount:          10,
		}

		mockWallet := entity.Wallet{
			Currency: "TRY",
			Balance:  0,
			UserId:   1,
		}

		mockExchange := entity.Exchange{
			Currency: "TRY",
			Rate:     1,
		}

		mockWalletRepo.On("Get", mock.Anything, mock.AnythingOfType("int64")).Return(mockWallet, nil).Once()

		mockExchangeRepo.On("GetByCurrency", mock.Anything, "TRY").Return(mockExchange, nil)

		//if is balance correct value return true
		mockWalletRepo.On("Update", mock.Anything, mock.MatchedBy(func(w entity.Wallet) bool {
			return w.Balance == 10
		})).Return(nil)

		mockTransactionRepo.On("Store", mock.Anything, mock.MatchedBy(func(t entity.Transaction) bool {
			return t.TransactionType == "withdraw"
		})).Return(nil)

		tuc := usecase.NewTransactionUseCase(mockTransactionRepo, mockWalletRepo, mockExchangeRepo)

		err := tuc.Post(context.Background(), mockTransaction)
		if err != nil {
			println(err)
		}

		assert.NoError(t, err)

		mockWalletRepo.AssertExpectations(t)
		mockExchangeRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})
}
