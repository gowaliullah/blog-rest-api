package databse

import "github.com/gowaliullah/blog-rest-api/models"

var PostList []models.Post

func Posts() {

	post1 := models.Post{ID: 1, Title: "The Future of AI in Web Development", Content: "A deep dive into how artificial intelligence is changing the landscape of web development, from automated code generation to personalized user experiences.", Images: []string{"ai-image-1.jpg", "ai-image-2.png"}, AuthorId: 101}
	post2 := models.Post{ID: 2, Title: "Healthy Eating on a Budget", Content: "Tips and tricks for making nutritious meals without breaking the bank. Includes a sample grocery list and five easy-to-follow recipes.", Images: []string{"food-photo-1.jpg"}, AuthorId: 102}
	post3 := models.Post{ID: 3, Title: "My First Solo Trip to Japan", Content: "An a-z guide to traveling through Japan alone. Covers everything from navigating the train system to finding the best ramen spots.", Images: []string{"japan-1.jpg", "japan-2.jpg", "japan-3.png"}, AuthorId: 101}
	post4 := models.Post{ID: 4, Title: "An Introduction to Rust Programming", Content: "This post will walk you through the basics of the Rust programming language, highlighting its key features and why it's gaining popularity.", Images: []string{}, AuthorId: 103}
	post5 := models.Post{ID: 5, Title: "Five Ways to Improve Your Home Workout", Content: "Whether you're a beginner or a seasoned pro, these five simple changes can help you get more out of your home fitness routine.", Images: []string{"workout-1.jpg"}, AuthorId: 104}

	PostList = append(PostList, post1, post2, post3, post4, post5)

}
