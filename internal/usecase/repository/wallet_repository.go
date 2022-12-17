package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	sq "github.com/Masterminds/squirrel"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/pkg/db"
)

// WalletRepo
type WalletRepository struct {
	*db.Postgres
}

// New Constractor
func NewWalletRepository(pg *db.Postgres) *WalletRepository {
	return &WalletRepository{pg}
}

// Get wallet

func (r *WalletRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Wallet, error) {
	//build sql string
	sql, args, err := r.Builder.
		Select("*").
		From("wallet").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil ,fmt.Errorf("WalletRepository - Get - r.Builder: %w", err)
	}

	//execute select query
	row := r.Db.QueryRowContext(ctx, sql, args...)
	if err != nil {
		return nil ,fmt.Errorf("WalletRepository - Get - r.Db.QueryRow: %w", err)
	}

	e := entity.Wallet{}

	err = row.Scan(&e.Currency, &e.Balance, &e.Id, &e.UserId)
	if err != nil {
		return nil, fmt.Errorf("WalletRepository - Get - rows.Scan: %w", err)
	}

	return &e, nil
}

// Insert wallet
func (r *WalletRepository) Store(ctx context.Context, t entity.Wallet) error {
	//build sql string
	sql, args, err := r.Builder.
		Insert("Wallet").
		Columns("id, balance, user_id, currency").
		Values(t.Id, t.Balance, t.Currency, t.UserId).
		ToSql()
	if err != nil {
		return fmt.Errorf("WalletRepository - Store - r.Builder: %w", err)
	}

	//execute insert command
	_, err = r.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("WalletRepository - Store - r.Db.Exec: %w", err)
	}

	return nil
}

func (r *WalletRepository) Update(ctx context.Context, t entity.Wallet) error {
	//build sql string
	sql, args, err := r.Builder.
		Update("Wallet").
		Set("balance, user_id, currency", &t).
		ToSql()
	if err != nil {
		return fmt.Errorf("WalletRepository - Update - r.Builder: %w", err)
	}

	//execute update command
	_, err = r.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("WalletRepository - Update - r.Db.Exec: %w", err)
	}

	return nil
}
