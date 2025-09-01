package util

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}
