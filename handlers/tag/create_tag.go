package tag

import (
	"encoding/json"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
	"github.com/gowaliullah/blog-rest-api/models"
)

func CreateTag(w http.ResponseWriter, r *http.Request) {
	var newTag models.Tag
	err := json.NewDecoder(r.Body).Decode(&newTag)
	if err != nil {
		http.Error(w, "please give valid JSON data", http.StatusBadRequest)
		return
	}

	newTag.ID = len(databse.TagList) + 1
	databse.TagList = append(databse.TagList, newTag)

	// set response type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// send the response
	json.NewEncoder(w).Encode(newTag)

}
