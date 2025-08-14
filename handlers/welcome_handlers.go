package handlers

import (
	"fmt"
	"net/http"
)

var WelcomeHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my blog server API, thanks for visiting")
}
