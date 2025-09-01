package users

import (
	"encoding/json"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/databse"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	encoder := json.NewEncoder(w)
	encoder.Encode(databse.UserList)
}
