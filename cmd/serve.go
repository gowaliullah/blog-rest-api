package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/handlers"
	"github.com/gowaliullah/blog-rest-api/handlers/category"
	"github.com/gowaliullah/blog-rest-api/handlers/comment"
	"github.com/gowaliullah/blog-rest-api/handlers/menu"
	"github.com/gowaliullah/blog-rest-api/handlers/posts"
	"github.com/gowaliullah/blog-rest-api/handlers/tag"
	"github.com/gowaliullah/blog-rest-api/handlers/users"
)

func Serve() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.WelcomeHandler)
	mux.HandleFunc("GET /about", handlers.AboutHandler)

	// users related routes
	mux.HandleFunc("POST /users", users.CreateUser)
	mux.HandleFunc("GET /users", users.GetUsers)
	mux.HandleFunc("GET /users/{id}", users.GetSingleUser)
	mux.HandleFunc("DELETE /users/{id}", users.DeleteUser)

	// users related routes
	mux.HandleFunc("POST /posts", posts.CreatePosts)
	mux.HandleFunc("GET /posts", users.GetUsers)
	mux.HandleFunc("GET /posts/{id}", users.GetSingleUser)

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

	fmt.Println("Server running on PORT: 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
