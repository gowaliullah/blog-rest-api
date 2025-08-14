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
	// convert the id
	id, err := strconv.Atoi(userId)

	// id the id is not valid send err
	if err != nil {
		http.Error(w, "Please send the valid user id", 400)
		return
	}

	for _, user := range databse.UserList {
		// check the user id
		if user.ID == id {
			// send the matched user
			w.WriteHeader(200)
			encoder := json.NewEncoder(w)
			encoder.Encode(user)
			return
		}
	}
	// do not matched the use
	w.WriteHeader(404)
	encoder := json.NewEncoder(w)
	encoder.Encode("User Not Found.....")
}
