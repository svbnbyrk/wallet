package usecase

import (
	"context"
	"fmt"

	"github.com/svbnbyrk/wallet/internal/entity"
)

type TransactionUseCase struct {
	tr TransactionRepository
	wr WalletRepository
}

func NewTransactionUsecase(t TransactionRepository, w WalletRepository) *TransactionUseCase {
	return &TransactionUseCase{
		tr: t,
		wr: w,
	}
}

// History - getting translate history from store.
func (uc *TransactionUseCase) History(ctx context.Context) ([]entity.Transaction, error) {
	transaction, err := uc.tr.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TransactionUseCase - History - s.repo.GetHistory: %w", err)
	}

	return transaction, nil
}

func (uc *TransactionUseCase) Post(ctx context.Context, u entity.Transaction) error {
	wallet, err := uc.wr.Get(ctx, u.WalletId)
	if err != nil {
		return fmt.Errorf("TransactionUseCase - Post - uc.wr.Get: %w", err)
	}

	var balance float64
	//to do translate currenct of amount
	var amount float64

	switch u.TransactionType {
	case "deposit":
		balance = wallet.Balance - amount
	case "withdrawal":
		balance = wallet.Balance + amount
	}

	wallet.Balance = balance

	err = uc.wr.Update(ctx, wallet)
	if err != nil {
		return fmt.Errorf("TransactionUseCase - Post - uc.wr.Update: %w", err)
	}

	transaction := entity.Transaction{
		WalletId:        u.WalletId,
		TransactionType: u.TransactionType,
		Currency:        u.Currency,
		Amount:          amount,
		Balance:         balance,
	}
	err = uc.tr.Store(ctx, transaction)
	if err != nil {
		return fmt.Errorf("TransactionUseCase - Post - uc.wr.Store: %w", err)
	}

	return nil
}
