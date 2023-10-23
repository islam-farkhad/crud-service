package postgres

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"

	"homework-3/internal/pkg/db"
)

// TestDB represents database for testing
type TestDB struct {
	DB db.DBops
	sync.Mutex
}

// NewFromEnv returns an instance of TestDB
func NewFromEnv() *TestDB {
	database, err := db.NewDB(context.Background(), db.GetDBConnectionString()) // тут должен передавать креды для тестовой бд
	if err != nil {
		panic(err)
	}
	return &TestDB{DB: database}
}

// SetUp clears locks and truncates test db
func (d *TestDB) SetUp(t *testing.T) {
	t.Helper()
	d.Lock()
	d.Truncate(context.Background())
}

// TearDown truncates db and unlocks it
func (d *TestDB) TearDown() {
	defer d.Unlock()
	d.Truncate(context.Background())
}

// Truncate clears test data base
func (d *TestDB) Truncate(ctx context.Context) {
	var tables []string
	err := d.DB.Select(ctx, &tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name != 'goose_db_version'")
	if err != nil {
		panic(err)
	}
	if len(tables) == 0 {
		panic("run migration please")
	}
	q := fmt.Sprintf("Truncate table %s", strings.Join(tables, ","))
	if _, err := d.DB.Exec(ctx, q); err != nil {
		panic(err)
	}
}
