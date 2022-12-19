package entity

import "context"

type Wallet struct {
	ID       int64   `json:"id"`
	UserId   int64   `json:"user_id"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

func NewWallet(userId int64, currency string, balance float64) Wallet {
	return Wallet{
		UserId:   userId,
		Currency: currency,
		Balance:  balance,
	}
}

// Wallet Usecase
type WalletUseCase interface {
	Store(ctx context.Context, w Wallet) error
	GetWalletsByUser(ctx context.Context, id int64) ([]Wallet, error)
}

// Wallet Repository
type WalletRepository interface {
	Get(ctx context.Context, id int64) (Wallet, error)
	Store(ctx context.Context, w Wallet) error
	Update(ctx context.Context, w Wallet) error
	GetWalletsByUser(ctx context.Context, id int64) ([]Wallet, error)
}
