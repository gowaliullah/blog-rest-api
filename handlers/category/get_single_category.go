package category

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowaliullah/blog-rest-api/databse"
)

func GetSingleCategory(w http.ResponseWriter, r *http.Request) {
	catId := r.PathValue("id")
	id, err := strconv.Atoi(catId)
	if err != nil {
		http.Error(w, "id not valid", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	for _, cat := range databse.CategoryList {
		if cat.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cat)
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Not found.....")
}
