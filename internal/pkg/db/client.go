package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NewDB is used to construct new connections pool
func NewDB(ctx context.Context, connectionString string) (*Database, error) {
	connectionsPool, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}
	return newDatabase(connectionsPool), nil
}
