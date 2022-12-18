package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/svbnbyrk/wallet/internal/entity"
)

type TransactionUseCase struct {
	tr TransactionRepository
	wr WalletRepository
	er ExchangeRepository
}

func NewTransactionUseCase(t TransactionRepository, w WalletRepository, e ExchangeRepository) *TransactionUseCase {
	return &TransactionUseCase{
		tr: t,
		wr: w,
		er: e,
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
	var balance float64
	//set amount
	amount := u.Amount

	wallet, err := uc.wr.Get(ctx, u.WalletId)
	if err != nil {
		return fmt.Errorf("TransactionUseCase - Post - uc.wr.Get: %w", err)
	}

	transactionRate, err := uc.er.GetByCurrency(ctx, u.Currency)
	if err != nil {
		return fmt.Errorf("TransactionUseCase - Post - uc.wr.Get: %w", err)
	}

	walletRate, err := uc.er.GetByCurrency(ctx, wallet.Currency)
	if err != nil {
		return fmt.Errorf("TransactionUseCase - Post - uc.wr.Get: %w", err)
	}

	if u.Currency != wallet.Currency {
		amount = amount * (walletRate.Rate / transactionRate.Rate)
	}

	switch u.TransactionType {
	case "deposit":
		balance = wallet.Balance - amount
	case "withdraw":
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
		Amount:          u.Amount,
		Balance:         balance,
		CreatedAt:       time.Now(),
	}
	err = uc.tr.Store(ctx, transaction)
	if err != nil {
		return fmt.Errorf("TransactionUseCase - Post - uc.wr.Store: %w", err)
	}

	return nil
}
