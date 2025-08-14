package databse

import "github.com/gowaliullah/blog-rest-api/models"

var CommentList []models.Comment

func Comments() {

	cmn1 := models.Comment{ID: 101, PostID: 1, AuthorID: 205, Content: "This is a great point about AI in web development! I'm curious about the specific tools you recommend for code generation."}
	cmn2 := models.Comment{ID: 102, PostID: 2, AuthorID: 201, Content: "I tried your recipe for lentil soup and it was a huge hit! Thanks for sharing this."}
	cmn3 := models.Comment{ID: 103, PostID: 1, AuthorID: 203, Content: "AI is moving so fast. It's both exciting and a little intimidating."}
	cmn4 := models.Comment{ID: 104, PostID: 3, AuthorID: 202, Content: "Your travel guide for Japan is a lifesaver. I'm heading there next month and this is exactly what I needed!"}
	cmn5 := models.Comment{ID: 105, PostID: 5, AuthorID: 204, Content: "Great tips! I've been struggling to stay motivated with my home workouts, and these ideas are really helpful."}

	CommentList = append(CommentList, cmn1, cmn2, cmn3, cmn4, cmn5)

}
