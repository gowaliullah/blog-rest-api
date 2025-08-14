package handlers

import (
	"fmt"
	"net/http"
)

var AboutHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Our purpose to service real world news supply")
}
