package databse

import "github.com/gowaliullah/blog-rest-api/models"

var MenuList []models.Menu

func Menus() {

	menu1 := models.Menu{ID: 1, Title: "Home", Url: "/", Slug: "home", ParentID: nil}
	menu2 := models.Menu{ID: 2, Title: "About Us", Url: "/about", Slug: "about-us", ParentID: nil}
	menu3 := models.Menu{ID: 3, Title: "Services", Url: "/services", Slug: "services", ParentID: nil}
	menu4 := models.Menu{ID: 4, Title: "Web Development", Url: "/services/web-dev", Slug: "web-development", ParentID: &[]int{3}[0]}
	menu5 := models.Menu{ID: 5, Title: "Contact", Url: "/contact", Slug: "contact", ParentID: nil}

	MenuList = append(MenuList, menu1, menu2, menu3, menu4, menu5)
}
