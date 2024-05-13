package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewConn() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), "postgres://root:0000@localhost:5432/aqary_task")
}
