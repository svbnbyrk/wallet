package repository

import (
	"context"
	"fmt"

	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/pkg/db"
)

// UserRepo
type UserRepository struct {
	*db.Postgres
}

// New Constractor
func NewUserRepository(pg *db.Postgres) *UserRepository {
	return &UserRepository{pg}
}

// Insert user
func (r *UserRepository) Store(ctx context.Context, u entity.User) error {
	//build sql string
	sql, args, err := r.Builder.
		Insert("users").
		Columns("email, name").
		Values(u.Email, u.Name).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepository - Store - r.Builder: %w", err)
	}

	//execute insert command
	_, err = r.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepository - Store - r.Db.Exec: %w", err)
	}

	return nil
}

// Update user
func (r *UserRepository) Update(ctx context.Context, u entity.User) error {
	//build sql string
	sql, args, err := r.Builder.
		Update("users").
		Set("email, name", &u).
		ToSql()
	if err != nil {
		return err
	}

	//execute update command
	_, err = r.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
