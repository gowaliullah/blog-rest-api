package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var config Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	DB          *sql.DB
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("fail to load .env variables:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("VERSION is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("SERVICE_NAME is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("HTTP_PORT is required")
		os.Exit(1)
	}

	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		log.Fatalf("DATABASE_URL is not loaded: %v", err)
		os.Exit(1)
	}

	// Postgres connection
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	fmt.Println("✅ Successfully connected to database!")

	config = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		DB:          db,
	}

}

func GetConfig() Config {
	loadConfig()
	return config
}

func ConnectDB() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// DB
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
