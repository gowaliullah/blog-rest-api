package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB() (*sql.DB, error) {
	connStr := "user=postgres password=postgres dbname=blogdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	fmt.Println("Successfully connected to db && server running on PORT: http://localhost:8080")
	return db, nil
}
