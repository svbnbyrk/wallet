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
	sq, args, err := r.Builder.
		Select("id, user_id, balance, currency").
		From("wallets").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("WalletRepository - Get - r.Builder: %w", err)
	}

	//execute select query
	row := r.Db.QueryRowContext(ctx, sq, args...)
	if err != nil {
		return entity.Wallet{}, err
	}

	e := entity.Wallet{}
	err = row.Scan(&e.ID, &e.UserId, &e.Balance, &e.Currency)

	if err != nil {
		return entity.Wallet{}, err
	}

	return e, nil
}

// Insert wallet
func (r *WalletRepository) Store(ctx context.Context, w entity.Wallet) error {
	//build sql string
	sql, args, err := r.Builder.
		Insert("wallets").
		Columns("balance, currency, user_id").
		Values(w.Balance, w.Currency, w.UserId).
		ToSql()
	if err != nil {
		return err
	}

	//execute insert command
	_, err = r.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("WalletRepository - Store - r.Db.Exec: %w", err)
	}

	return nil
}

func (r *WalletRepository) Update(ctx context.Context, w entity.Wallet) error {
	//build sql string
	sql, args, err := r.Builder.
		Update("wallets").
		Set("balance", w.Balance).
		Set("user_id", w.UserId).
		Set("currency", w.Currency).
		Where(sq.Eq{"id": w.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("WalletRepository - Update - r.Builder: %w", err)
	}

	//execute update command
	_, err = r.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("WalletRepository - Update - r.Db.ExecContext: %w", err)
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
