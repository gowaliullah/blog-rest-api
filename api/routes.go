package api

import (
	"database/sql"
	"net/http"
)

// SetupRoutes registers all the API endpoints with the router
func SetupRoutes(router *http.ServeMux, db *sql.DB) {
	router.HandleFunc("/menus", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			id := r.URL.Query().Get("id")
			if id != "" {
				GetSingleCategiry(db, w, r, id)
			} else {
				GetAllCategories(db, w, r)
			}
		// case "POST":
		// 	CreateCategory(db, w, r)
		// case "PUT":
		// 	UpdateCategory(db, w, r)
		// case "DELETE":
		// 	DeleteCategory(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
