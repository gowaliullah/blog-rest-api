package main

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
)

var welcomeHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my blog server API, thanks for visiting")
}
var aboutHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Our purpose to service real world news supply")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", welcomeHandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on PORT: 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

func init() {
	databse.Users()
}
