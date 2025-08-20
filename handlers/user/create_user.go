package users

import (
	"encoding/json"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
	"github.com/gowaliullah/blog-rest-api/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	// define variable for newUser
	var newUser models.User

	// decode user data
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "please give valid JSON data", http.StatusBadRequest)
		return
	}

	// manually increament user ID
	newUser.ID = len(databse.UserList) + 1

	// append new user data
	databse.UserList = append(databse.UserList, newUser)

	// set response type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// send the response
	json.NewEncoder(w).Encode(newUser)

}
