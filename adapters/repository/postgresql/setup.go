package postgresql

import (
	"context"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/filipeandrade6/fiap-pedeai-clientes/adapters/repository/postgresql/db"
)

type Config struct {
	Host       string
	Port       string
	User       string
	Password   string
	Name       string
	DisableTLS bool
}

type Repository struct {
	db *db.Queries
}

func New(ctx context.Context, cfg Config) (*Repository, error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	pgxPool, err := pgxpool.New(ctx, u.String())
	if err != nil {
		return nil, err
	}

	repo := db.New(pgxPool)

	return &Repository{repo}, nil
}
