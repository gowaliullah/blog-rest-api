package comment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowaliullah/blog-rest-api/databse"
)

func GetSingleComment(w http.ResponseWriter, r *http.Request) {
	commentId := r.PathValue("id")
	id, err := strconv.Atoi(commentId)
	if err != nil {
		http.Error(w, "id not valid", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	for _, comment := range databse.CommentList {
		if comment.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(comment)

		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Comment not found")
}
