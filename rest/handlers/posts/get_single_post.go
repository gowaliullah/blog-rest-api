package posts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowaliullah/blog-rest-api/databse"
)

func GetSinglePost(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")
	id, err := strconv.Atoi(postId)

	if err != nil {
		http.Error(w, "Please send the valid post id", http.StatusBadRequest)
		return
	}

	// set response content type
	w.Header().Set("Content-Type", "application/json")

	for _, post := range databse.PostList {
		if post.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("User Not Found.....")
}
