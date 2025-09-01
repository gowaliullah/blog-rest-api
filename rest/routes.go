package rest

import (
	"net/http"

	"github.com/gowaliullah/blog-rest-api/rest/handlers/menu"
)

func initRoutes(mux *http.ServeMux) {
	// users related routes
	mux.HandleFunc("POST /users", menu.CreateMenu)
	mux.HandleFunc("GET /users", menu.GetAllMenus)
}
