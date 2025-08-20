package category

import (
	"encoding/json"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
	"github.com/gowaliullah/blog-rest-api/models"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCat models.Category

	err := json.NewDecoder(r.Body).Decode(&newCat)
	if err != nil {
		http.Error(w, "please give valid JSON data", http.StatusBadRequest)
		return
	}

	newCat.ID = len(databse.CategoryList) + 1
	databse.CategoryList = append(databse.CategoryList, newCat)

	// set response type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCat)

}
