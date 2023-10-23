//go:build integration

package tests

import (
	"homework-3/internal/crud"
	"homework-3/tests/app"
	"homework-3/tests/postgres"
)

var (
	database *postgres.TestDB
	testApp  crud.App
)

func init() {
	database = postgres.NewFromEnv()
	testApp = app.NewTestApp(database.DB)
}
