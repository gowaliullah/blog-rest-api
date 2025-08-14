package main

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
	"github.com/gowaliullah/blog-rest-api/handlers"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.WelcomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)

	fmt.Println("Server running on PORT: 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

func init() {
	databse.Users()
	databse.Posts()
	databse.Comments()
}
