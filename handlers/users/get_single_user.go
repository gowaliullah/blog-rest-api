package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowaliullah/blog-rest-api/databse"
)

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	// get id from the header path
	userId := r.PathValue("id")

	// convert the id to int
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Please send the valid user id", http.StatusBadRequest)
		return
	}

	// set response content type
	w.Header().Set("Content-Type", "application/json")

	// search user by ID
	for _, user := range databse.UserList {
		if user.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	// if user not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("User Not Found.....")
}
