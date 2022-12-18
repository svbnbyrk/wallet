package usecase

import (
	"context"

	"github.com/svbnbyrk/wallet/internal/entity"
)

//go:generate mockgen -source=interface.go -destination=./mocks_test.go -package=usecase_test

type (
	// Transaction Usecase
	TransactionUsecase interface {
		History(context.Context) ([]entity.Transaction, error)
		Post(context.Context, entity.Transaction) error
	}

	// TransactionRepo
	TransactionRepository interface {
		Store(context.Context, entity.Transaction) error
		GetHistory(context.Context) ([]entity.Transaction, error)
	}

	// Wallet Usecase
	WalletUsecase interface {
		Store(context.Context, entity.Wallet) error
		GetWalletsbyUser(context.Context, int64) ([]entity.Wallet, error)
	}

	// WalletRepo
	WalletRepository interface {
		Get(context.Context, int64) (entity.Wallet, error)
		Store(context.Context, entity.Wallet) error
		Update(context.Context, entity.Wallet) error
		GetbyUserId(context.Context, int64) ([]entity.Wallet, error)
	}

	// User Usecase
	UserUsecase interface {
		Store(context.Context, entity.User) error
	}

	// UserRepo
	UserRepository interface {
		Store(context.Context, entity.User) error
		Update(context.Context, entity.User) error
	}
)
