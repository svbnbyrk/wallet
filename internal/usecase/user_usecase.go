package usecase

import (
	"context"
	"fmt"

	"github.com/svbnbyrk/wallet/internal/entity"
)

type UserUseCase struct {
	repo UserRepository
}

func NewUserUsecase(r UserRepository) *UserUseCase {
	return &UserUseCase{
		repo:   r,
	}
}

// Store - Insert user
func (uc *UserUseCase) Store(ctx context.Context, u entity.User )  error {
	err := uc.repo.Store(ctx, u)
	if err != nil {
		return fmt.Errorf("UserUseCase - Post - s.repo.Store: %w", err)
	}

	return nil
}
