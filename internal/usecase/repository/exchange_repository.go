package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/svbnbyrk/wallet/internal/entity"
	"github.com/svbnbyrk/wallet/pkg/db"
)

// ExchangeRepo
type ExchangeRepository struct {
	*db.Postgres
}

// New Constractor
func NewExchangeRepository(pg *db.Postgres) *ExchangeRepository {
	return &ExchangeRepository{pg}
}

// Get Exchange

func (r *ExchangeRepository) GetByCurrency(ctx context.Context, currency string) (entity.Exchange, error) {
	//build sql string
	sql, args, err := r.Builder.
		Select("rate, currency").
		From("exchanges").
		Where(sq.Eq{"currency": currency}).
		ToSql()
	if err != nil {
		return entity.Exchange{}, fmt.Errorf("ExchangeRepository - Get - r.Builder: %w", err)
	}

	//execute select query
	row := r.Db.QueryRowContext(ctx, sql, args...)
	if err != nil {
		return entity.Exchange{}, fmt.Errorf("ExchangeRepository - Get - r.Db.QueryRow: %w", err)
	}

	e := entity.Exchange{}

	err = row.Scan(&e.Rate, &e.Currency)
	if err != nil {
		return entity.Exchange{}, fmt.Errorf("ExchangeRepository - Get - rows.Scan: %w", err)
	}

	return e, nil
}
