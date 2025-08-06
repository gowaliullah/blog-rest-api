package main

import (
	"log"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/api"
	"github.com/gowaliullah/blog-rest-api/config"
)

func main() {
	// Initialize the database connection
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the router
	router := http.NewServeMux()
	api.SetupRoutes(router, db)

	// fmt.Println("Server is running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
