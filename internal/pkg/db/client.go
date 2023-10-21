package db

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	testHost     = "localhost"
	testPortStr  = "5432"
	testUser     = "test"
	testPassword = "test"
	testDBName   = "test"
)

// NewDB is used to construct new connections pool
func NewDB(ctx context.Context) (*Database, error) {
	connectionsPool, err := pgxpool.Connect(ctx, getDBConnectionString())
	if err != nil {
		return nil, err
	}
	return newDatabase(connectionsPool), nil
}

func getDBConnectionString() string {
	portStr := os.Getenv("port")
	if portStr == "" {
		portStr = testPortStr
	}
	port, _ := strconv.Atoi(portStr)

	addr := os.Getenv("host")
	if addr == "" {
		addr = testHost
	}

	user := os.Getenv("user")
	if user == "" {
		user = testUser
	}

	password := os.Getenv("password")
	if password == "" {
		password = testPassword
	}

	dbName := os.Getenv("dbname")
	if dbName == "" {
		dbName = testDBName
	}

	config := Config{
		Addr:     addr,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Addr, config.Port, config.User, config.Password, config.DBName)
}
