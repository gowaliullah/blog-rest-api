package rest

import (
	"net/http"

	"github.com/gowaliullah/blog-rest-api/rest/handlers/users"
)

func initRoutes(mux *http.ServeMux) {
	// users related routes
	// mux.HandleFunc("GET /users", users.GetAllUsers)
	mux.HandleFunc("POST /users", users.CreateUser)
}
