package usecase

import (
	"context"
	"fmt"

	"github.com/svbnbyrk/wallet/internal/entity"
)

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(r UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

// Store user
func (uc *UserUseCase) Store(ctx context.Context, u entity.User) error {
	err := uc.repo.Store(ctx, u)
	if err != nil {
		return fmt.Errorf("UserUseCase - Post - s.repo.Store: %w", err)
	}

	return nil
}
