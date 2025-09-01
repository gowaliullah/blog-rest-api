package menu

import (
	"encoding/json"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
	"github.com/gowaliullah/blog-rest-api/models"
)

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	var newMenu models.Menu

	err := json.NewDecoder(r.Body).Decode(&newMenu)
	if err != nil {
		http.Error(w, "please give valid JSON data", http.StatusBadRequest)
		return
	}

	newMenu.ID = len(databse.MenuList) + 1
	databse.MenuList = append(databse.MenuList, newMenu)

	// set response type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMenu)

}
