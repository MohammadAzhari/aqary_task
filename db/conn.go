package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPool() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), "postgres://root:0000@localhost:5432/aqary_task")
	pool.Config().MaxConns = 20
	return pool, err
}
