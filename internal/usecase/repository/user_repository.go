package repository

import (
	"context"
	"fmt"

	"github.com/svbnbyrk/wallet/pkg/db"
	"github.com/svbnbyrk/wallet/internal/entity"
)

const _defaultEntityCap = 64

// UserRepo
type UserRepository struct {
	*db.Postgres
}

// New Constractor
func NewUserRepository(pg *db.Postgres) *UserRepository {
	return &UserRepository{pg}
}

// Insert user  
func (r *UserRepository) Store(ctx context.Context, t entity.User) error {
	//build sql string 
	sql, args, err := r.Builder.
		Insert("User").
		Columns("id, email, name").
		Values(t.Id, t.Email, t.Name).
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
func (r *UserRepository) Update(ctx context.Context, t entity.User) error {
	//build sql string 
	sql, args, err := r.Builder.
		Update("User").
		Set("email, name",&t).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepository - Update - r.Builder: %w", err)
	}

	//execute update command
	_, err = r.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepository - Update - r.Db.Exec: %w", err)
	}

	return nil
}


