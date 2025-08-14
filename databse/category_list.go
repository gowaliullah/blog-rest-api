package databse

import "github.com/gowaliullah/blog-rest-api/models"

var CategoryList []models.Category

func Categories() {
	cat1 := models.Category{ID: 1, Name: "Technology", Slug: "technology"}
	cat2 := models.Category{ID: 2, Name: "Travel", Slug: "travel"}
	cat3 := models.Category{ID: 3, Name: "Food & Cooking", Slug: "food-cooking"}
	cat4 := models.Category{ID: 4, Name: "Health & Fitness", Slug: "health-fitness"}
	cat5 := models.Category{ID: 5, Name: "Finance", Slug: "finance"}

	CategoryList = append(CategoryList, cat1, cat2, cat3, cat4, cat5)
}
