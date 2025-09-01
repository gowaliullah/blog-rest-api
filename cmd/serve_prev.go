package cmd

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gowaliullah/blog-rest-api/config"
	"github.com/gowaliullah/blog-rest-api/rest/handlers"
	"github.com/gowaliullah/blog-rest-api/rest/handlers/category"
	"github.com/gowaliullah/blog-rest-api/rest/handlers/comment"
	"github.com/gowaliullah/blog-rest-api/rest/handlers/menu"
	"github.com/gowaliullah/blog-rest-api/rest/handlers/tag"
	// "github.com/gowaliullah/blog-rest-api/handlers/users"
)

func Serve() {

	mux := http.NewServeMux()
	config.ConnectDB()

	mux.HandleFunc("/", handlers.WelcomeHandler)

	// users related routes
	// mux.HandleFunc("POST /users", users.CreateUser)
	// mux.HandleFunc("GET /users", users.GetAllUsers())

	// mux.HandleFunc("GET /users/{id}", users.GetSingleUser)
	// mux.HandleFunc("DELETE /users/{id}", users.DeleteUser)

	// // users related routes
	// mux.HandleFunc("POST /posts", posts.CreatePosts)
	// mux.HandleFunc("GET /posts", users.GetUsers)
	// mux.HandleFunc("GET /posts/{id}", users.GetSingleUser)

	// menus related routes
	mux.HandleFunc("POST /menus", menu.CreateMenu)
	mux.HandleFunc("GET /menus", menu.GetAllMenus)
	mux.HandleFunc("GET /menus/{id}", menu.GetSingleMenu)

	// comments related routes
	mux.HandleFunc("POST /comments", comment.CreateComment)
	mux.HandleFunc("GET /comments", comment.GetAllComments)
	mux.HandleFunc("GET /comments/{id}", menu.GetSingleMenu)

	// tags related routes
	mux.HandleFunc("POST /tags", tag.CreateTag)
	mux.HandleFunc("GET /tags", tag.GetAllTags)
	mux.HandleFunc("GET /tags/{id}", tag.GetSingleTag)

	// category related routes
	mux.HandleFunc("POST /categories", category.CreateCategory)
	mux.HandleFunc("GET /categories", category.GetAllCategories)
	mux.HandleFunc("GET /categories/{id}", category.GetSingleCategory)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
