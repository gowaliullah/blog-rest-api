package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gowaliullah/blog-rest-api/databse"
	"github.com/gowaliullah/blog-rest-api/models"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "this is not valid ID", http.StatusBadRequest)
	}

	var updatedUser []models.User

	for _, user := range databse.UserList {
		if user.ID != id {
			updatedUser = append(updatedUser, user)

			databse.UserList = updatedUser

			// set response type
			// w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		}

		// send the response
		// json.NewEncoder(w).Encode("user deleted successfully...")
	}

	// if user not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("User Not Found.....")

}
