package usecase

import (
	"context"
	"time"

	"github.com/svbnbyrk/wallet/internal/entity"
)

type TransactionUseCase struct {
	tr entity.TransactionRepository
	wr entity.WalletRepository
	er entity.ExchangeRepository
}

func NewTransactionUseCase(t entity.TransactionRepository, w entity.WalletRepository, e entity.ExchangeRepository) *TransactionUseCase {
	return &TransactionUseCase{
		tr: t,
		wr: w,
		er: e,
	}
}

// History - getting translate history from store.
func (uc *TransactionUseCase) History(ctx context.Context) ([]entity.Transaction, error) {
	transaction, err := uc.tr.GetHistory(ctx)
	if transaction == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (uc *TransactionUseCase) Post(ctx context.Context, u entity.Transaction) error {
	var balance float64

	//set amount
	amount := u.Amount

	wallet, err := uc.wr.Get(ctx, u.WalletId)
	if err != nil {
		return err
	}

	transactionRate, err := uc.er.GetByCurrency(ctx, u.Currency)
	if err != nil {
		return err
	}

	walletRate, err := uc.er.GetByCurrency(ctx, wallet.Currency)
	if err != nil {
		return err
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
		return err
	}

	transaction := entity.Transaction{
		WalletId:        u.WalletId,
		TransactionType: u.TransactionType,
		Currency:        u.Currency,
		Amount:          u.Amount,
		CreatedAt:       time.Now(),
	}

	err = uc.tr.Store(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
