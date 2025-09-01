package menu

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowaliullah/blog-rest-api/databse"
)

func GetSingleMenu(w http.ResponseWriter, r *http.Request) {
	menuId := r.PathValue("id")
	id, err := strconv.Atoi(menuId)
	if err != nil {
		http.Error(w, "id not valid", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	for _, m := range databse.MenuList {
		if m.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(m)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Menu not found...")
}
