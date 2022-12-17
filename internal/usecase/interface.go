package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/svbnbyrk/wallet/internal/entity"
)

//go:generate mockgen -source=interface.go -destination=./mocks/mocks_test.go -package=mocks_test

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
		Store(context.Context,entity.Wallet ) 
	}

	// WalletRepo
	WalletRepository interface {
		Get(context.Context, uuid.UUID) (*entity.Wallet, error)
		Store(context.Context, entity.Wallet) error
		Update(context.Context, entity.Wallet) error
	}

	// User Usecase
	UserUsecase interface {
		Store(context.Context, entity.User ) 
	}

	// UserRepo
	UserRepository interface {
		Store(context.Context, entity.User) error
		Update(context.Context, entity.User)
	}
)
