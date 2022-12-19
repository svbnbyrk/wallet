package usecase

import (
	"context"

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
		return err
	}

	return nil
}
