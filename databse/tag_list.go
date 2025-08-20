package databse

import "github.com/gowaliullah/blog-rest-api/models"

var TagList []models.Tag

func Tags() {

	tag1 := models.Tag{ID: 1, Name: "javascript", Slug: "javascript"}
	tag2 := models.Tag{ID: 2, Name: "travel-hacks", Slug: "travel-hacks"}
	tag3 := models.Tag{ID: 3, Name: "budget-recipes", Slug: "budget-recipes"}
	tag4 := models.Tag{ID: 4, Name: "machine-learning", Slug: "machine-learning"}
	tag5 := models.Tag{ID: 5, Name: "home-fitness", Slug: "home-fitness"}

	TagList = append(TagList, tag1, tag2, tag3, tag4, tag5)

}
