package comment

import (
	"encoding/json"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
)

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(databse.CommentList)
}
