package tag

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowaliullah/blog-rest-api/databse"
)

func GetSingleTag(w http.ResponseWriter, r *http.Request) {
	tagId := r.PathValue("id")
	id, err := strconv.Atoi(tagId)

	if err != nil {
		http.Error(w, "Please send the valid post id", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, tag := range databse.TagList {
		if tag.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tag)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("User Not Found.....")
}
