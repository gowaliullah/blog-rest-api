package comment

import (
	"encoding/json"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
	"github.com/gowaliullah/blog-rest-api/models"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comment
	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		http.Error(w, "please give valid JSON data", http.StatusBadRequest)
		return
	}

	newComment.ID = len(databse.CommentList) + 1
	databse.CommentList = append(databse.CommentList, newComment)

	// set response type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// send the response
	json.NewEncoder(w).Encode(newComment)
}
