package config

import (
	"fmt"
	"os"

	"github.com/gowaliullah/blog-rest-api/models"
	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables into a Config struct
func LoadConfig() (*models.Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &models.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	return cfg, nil
}
