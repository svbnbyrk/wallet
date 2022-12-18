package repository

import (
	"context"
	"fmt"

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

func (r *WalletRepository) Get(ctx context.Context, id int64) (entity.Wallet, error) {
	//build sql string
	sql, args, err := r.Builder.
		Select("*").
		From("wallet").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("WalletRepository - Get - r.Builder: %w", err)
	}

	//execute select query
	row := r.Db.QueryRowContext(ctx, sql, args...)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("WalletRepository - Get - r.Db.QueryRow: %w", err)
	}

	e := entity.Wallet{}

	err = row.Scan(&e.Currency, &e.Balance, &e.ID, &e.UserId)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("WalletRepository - Get - rows.Scan: %w", err)
	}

	return e, nil
}

// Insert wallet
func (r *WalletRepository) Store(ctx context.Context, t entity.Wallet) error {
	//build sql string
	sql, args, err := r.Builder.
		Insert("wallets").
		Columns("balance, currency, user_id").
		Values(t.Balance, t.Currency, t.UserId).
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
		Update("wallets").
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

func (r *WalletRepository) GetbyUserId(ctx context.Context, id int64) ([]entity.Wallet, error) {
	//build sql string
	sql, args, err := r.Builder.
		Select("id, balance, currency, user_id").
		From("wallets").
		Where(sq.Eq{"user_id": id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("WalletRepository - GetbyUserId - r.Builder: %w", err)
	}

	//execute select query
	rows, err := r.Db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("WalletRepository - GetbyUserId - r.Db.QueryContext: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Wallet, 0, 64)

	//fill rows to entity
	for rows.Next() {
		e := entity.Wallet{}

		err = rows.Scan(&e.ID, &e.Balance, &e.Currency, &e.UserId)
		if err != nil {
			return nil, fmt.Errorf("WalletRepository - GetbyUserId - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}
