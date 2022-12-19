package repository

import (
	"context"
	"fmt"

	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/pkg/db"
)

// TransactionRepo
type TransactionRepository struct {
	*db.Postgres
}

// New Constractor
func NewTransactionRepository(pg *db.Postgres) *TransactionRepository {
	return &TransactionRepository{pg}
}

// GetHistory
func (r *TransactionRepository) GetHistory(ctx context.Context) ([]entity.Transaction, error) {
	//build sql string
	sql, _, err := r.Builder.
		Select("id ,currency, transaction_type, wallet_id, amount, created_at").
		From("transactions").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("TransactionRepo - GetHistory - r.Builder: %w", err)
	}

	//execute sql query
	rows, err := r.Db.QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TransactionRepo - GetHistory - r.Db.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Transaction, 0, 64)

	//fill rows to entity
	for rows.Next() {
		e := entity.Transaction{}

		err = rows.Scan(&e.ID, &e.Currency, &e.TransactionType, &e.WalletId, &e.Amount, &e.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("TransactionRepo - GetHistory - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// Store -.
func (r *TransactionRepository) Store(ctx context.Context, t entity.Transaction) error {
	//build sql string
	sql, args, err := r.Builder.
		Insert("transactions").
		Columns("currency, transaction_type, wallet_id, amount , created_at").
		Values(t.Currency, t.TransactionType, t.WalletId, t.Amount, t.CreatedAt).
		ToSql()
	if err != nil {
		return fmt.Errorf("TransactionRepo - Store - r.Builder: %w", err)
	}

	//execute insert command
	_, err = r.Db.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("TransactionRepo - Store - r.Db.Exec: %w", err)
	}

	return nil
}
