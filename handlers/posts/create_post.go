package posts

import (
	"encoding/json"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
	"github.com/gowaliullah/blog-rest-api/models"
)

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post

	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "please give valid JSON data", http.StatusBadRequest)
		return
	}

	newPost.ID = len(databse.PostList) + 1

	databse.PostList = append(databse.PostList, newPost)

	// set response type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// send the response
	json.NewEncoder(w).Encode(newPost)

}
