package posts

import (
	"encoding/json"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(databse.PostList)
}
