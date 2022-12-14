package entity

import (
	"context"
	"time"
)

type TransactionType string

const (
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
)

type Transaction struct {
	//TODO ID
	ID              int64           `json:"id"`
	WalletId        int64           `json:"wallet_id" binding:"required"`
	TransactionType TransactionType `json:"transactionType" binding:"required,oneof=deposit withdraw"`
	Amount          float64         `json:"amount" binding:"required"`
	Currency        string          `json:"currency" binding:"required,iso4217"`
	CreatedAt       time.Time       `json:"created_at"`
}

func NewTransaction(walletId int64, transactionType TransactionType, currency string, amount float64) Transaction {
	return Transaction{
		WalletId:        walletId,
		TransactionType: transactionType,
		Currency:        currency,
		Amount:          amount,
	}
}

type TransactionUseCase interface {
	History(ctx context.Context) ([]Transaction, error)
	Post(ctx context.Context, t Transaction) error
}

// Transaction Repository
type TransactionRepository interface {
	Store(ctx context.Context, t Transaction) error
	GetHistory(ctx context.Context) ([]Transaction, error)
}
