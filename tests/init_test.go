//go:build integration

package tests

import "homework-3/tests/postgres"

var (
	database *postgres.TestDB
)

func init() {
	database = postgres.NewFromEnv()
}
