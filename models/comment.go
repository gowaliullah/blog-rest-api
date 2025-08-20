package models

type Comment struct {
	ID       int    `json:"_id"`
	PostID   int    `json:"post_id"`
	AuthorID int    `json:"author_id"`
	Content  string `json:"content"`
}
