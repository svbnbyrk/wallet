package usecase

import (
	"context"

	"github.com/svbnbyrk/wallet/internal/entity"
)

//go:generate mockgen mockery --all --keeptree

type (
	// Transaction Usecase
	Transaction interface {
		History(ctx context.Context) ([]entity.Transaction, error)
		Post(ctx context.Context, t entity.Transaction) error
	}

	// Transaction Repository
	TransactionRepository interface {
		Store(ctx context.Context, t entity.Transaction) error
		GetHistory(ctx context.Context) ([]entity.Transaction, error)
	}

	// Wallet Usecase
	Wallet interface {
		Store(ctx context.Context, w entity.Wallet) error
		GetWalletsbyUser(ctx context.Context, id int64) ([]entity.Wallet, error)
	}

	// Wallet Repository
	WalletRepository interface {
		Get(ctx context.Context, id int64) (entity.Wallet, error)
		Store(ctx context.Context, w entity.Wallet) error
		Update(ctx context.Context, w entity.Wallet) error
		GetbyUserId(ctx context.Context, id int64) ([]entity.Wallet, error)
	}

	// User Usecase
	User interface {
		Store(ctx context.Context, u entity.User) error
	}

	// User Repository
	UserRepository interface {
		Store(ctx context.Context, u entity.User) error
		Update(ctx context.Context, u entity.User) error
	}

	// Exchange repository
	ExchangeRepository interface {
		GetByCurrency(ctx context.Context, currency string) (entity.Exchange, error)
	}
)
