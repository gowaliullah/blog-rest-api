package cmd

import (
	"github.com/gowaliullah/blog-rest-api/config"
	"github.com/gowaliullah/blog-rest-api/rest"
)

func Serve() {

	cnf := config.GetConfig()

	rest.Start(cnf)

}
