//go:build integration

package tests

import (
	"crud-service/internal/handlers"
	"crud-service/tests/app"
	"crud-service/tests/postgres"
)

var (
	database *postgres.TestDB
	testApp  handlers.App
)

func init() {
	database = postgres.NewFromEnv()
	testApp = app.NewTestApp(database.DB)
}
