package usecase

import (
	"context"

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
func (uc *WalletUseCase) Store(ctx context.Context, w entity.Wallet) error {
	err := uc.repo.Store(ctx, w)
	if err != nil {
		return err
	}

	return nil
}

func (uc *WalletUseCase) GetWalletsbyUser(ctx context.Context, id int64) ([]entity.Wallet, error) {
	wallets, err := uc.repo.GetbyUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}
