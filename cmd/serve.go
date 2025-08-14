package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/handlers"
	"github.com/gowaliullah/blog-rest-api/handlers/users"
)

func Serve() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.WelcomeHandler)
	mux.HandleFunc("GET /about", handlers.AboutHandler)
	mux.HandleFunc("POST /users", users.CreateUser)
	mux.HandleFunc("GET /users", users.GetUsers)
	mux.HandleFunc("GET /users/{id}", users.GetSingleUser)

	fmt.Println("Server running on PORT: 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
