package usecase

import (
	"context"
	"fmt"

	"github.com/svbnbyrk/wallet/internal/entity"
)

type WalletUseCase struct {
	repo WalletRepository
}

func NewWalletUseCase(r WalletRepository) *WalletUseCase {
	return &WalletUseCase{
		repo: r,
	}
}

// Store - Insert Wallet
func (uc *WalletUseCase) Store(ctx context.Context, u entity.Wallet) error {
	err := uc.repo.Store(ctx, u)
	if err != nil {
		return fmt.Errorf("WalletUseCase - Post - s.repo.Store: %w", err)
	}

	return nil
}

func (uc *WalletUseCase) GetWalletsbyUser(ctx context.Context, id int64) ([]entity.Wallet, error) {
	wallets, err := uc.repo.GetbyUserId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("TransactionUseCase - GetWalletsbyUser - s.repo.GetHistory: %w", err)
	}

	return wallets, nil
}
