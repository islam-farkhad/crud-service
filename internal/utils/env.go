package utils

import (
	"crud-service/internal/pkg/db"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func findEnvFile() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, ".env")); err != nil {
			parentDir := filepath.Dir(currentDir)

			if parentDir == currentDir {
				break
			}
			currentDir = parentDir
			continue
		}

		return filepath.Join(currentDir, ".env"), nil
	}

	return "", fmt.Errorf(".env file not found")
}

func init() {
	envFilePath, err := findEnvFile()
	if err != nil {
		log.Fatal(".env file not found")
	}

	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %+v", err)
	}
}

// GetAPIPort retrieves the API port from the environment variable "HTTP_PORT".
func GetAPIPort() string {
	apiPort := os.Getenv("HTTP_PORT")
	if apiPort == "" {
		log.Fatal("HTTP_PORT environment variable is not set")
	}
	return ":" + apiPort
}

// GetEnvDBConnectionConfig builds db.Config instance based on environment variables
func GetEnvDBConnectionConfig() *db.Config {
	return &db.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	}
}
