package utils

import (
	"context"
	"crud-service/internal/pkg/db"
	"fmt"
	"log"
)

// ConnectDB connects to db and returns Database instance
func ConnectDB(ctx context.Context) *db.Database {
	database, err := db.NewDB(ctx, MakeDBConnStr(GetEnvDBConnectionConfig()))
	if err != nil {
		log.Fatal(err)
	}
	return database
}

// MakeDBConnStr construct db connection string based on passed db.Config instance
func MakeDBConnStr(config *db.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
}
