package repository

import (
	"context"
	"database/sql"

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
	//build sqlQuery string
	sqlQuery, args, err := r.Builder.
		Select("rate, currency").
		From("exchanges").
		Where(sq.Eq{"currency": currency}).
		ToSql()
	if err != nil {
		return entity.Exchange{}, err
	}

	//execute select query
	row := r.Db.QueryRowContext(ctx, sqlQuery, args...)
	if err != nil {
		return entity.Exchange{}, err
	}

	e := entity.Exchange{}

	err = row.Scan(&e.Rate, &e.Currency)
	if err != nil && err != sql.ErrNoRows {
		return entity.Exchange{}, err
	}

	return e, nil
}
