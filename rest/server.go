package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gowaliullah/blog-rest-api/config"
)

func Start(cnf config.Config) {
	mux := http.NewServeMux()

	initRoutes(mux)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("server running on PORT:", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}
