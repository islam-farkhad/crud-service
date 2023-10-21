//go:generate mockgen -source ./database.go -destination=./mocks/database.go -package=mock_database
package db

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PGX interface {
	DBops
	BeginTX(ctx context.Context, options *pgx.TxOptions)
}

type DBops interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	GetConnectionsPool(_ context.Context) *pgxpool.Pool
}

// Database struct is used to store connection with database
type Database struct {
	connectionsPool *pgxpool.Pool
}

func newDatabase(connectionsPool *pgxpool.Pool) *Database {
	return &Database{
		connectionsPool: connectionsPool,
	}
}

// GetConnectionsPool returns the PostgreSQL connection pool associated with the database instance.
// It provides access to the underlying connection pool for executing queries and transactions.
func (db Database) GetConnectionsPool(_ context.Context) *pgxpool.Pool {
	return db.connectionsPool
}

// Get performs a SQL SELECT query using the given query string and optional arguments.
// It retrieves a single row result and populates the provided 'dest' struct with the data.
// The function returns an error if the query execution or result scanning fails.
func (db Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.connectionsPool, dest, query, args...)
}

// Select performs a SQL SELECT query using the given query string and optional arguments.
// It populates the provided 'dest' slice or struct with the results.
// The function returns an error if the query execution or result scanning fails.
func (db Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.connectionsPool, dest, query, args...)
}

// Exec executes the given SQL query with optional arguments and returns the result.
// It is used for SQL queries that don't return rows, such as INSERT, UPDATE, DELETE, etc.
// The pgconn.CommandTag and an error are returned.
func (db Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.connectionsPool.Exec(ctx, query, args...)
}

// ExecQueryRow executes a query that is expected to return at most one row and returns a pgx.Row object.
func (db Database) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.connectionsPool.QueryRow(ctx, query, args...)
}
