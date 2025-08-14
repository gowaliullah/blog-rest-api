package main

import (
	"github.com/gowaliullah/blog-rest-api/cmd"
	databse "github.com/gowaliullah/blog-rest-api/databse"
)

func main() {
	cmd.Serve()
}

func init() {
	databse.Users()
	databse.Posts()
	databse.Comments()
	databse.Menus()
	databse.Tags()
	databse.Categories()
}
