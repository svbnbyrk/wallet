package db

import (
	"embed"
	"fmt"

	"database/sql"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migration/*.sql
var embedMigrations embed.FS

// Postgres
type Postgres struct {
	Builder squirrel.StatementBuilderType
	Db      *sql.DB
}

// Db connection
func New(url string) (*Postgres, error) {
	var err error
	pg := &Postgres{}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	pg.Db, err = sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - sql.Open %w", err)
	}

	goose.SetBaseFS(embedMigrations)

	if err = goose.SetDialect("postgres"); err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - goose.SetDialect %w", err)
	}

	if err := goose.Up(pg.Db, "migration"); err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - goose.Up %w", err)
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Db != nil {
		p.Db.Close()
	}
}
