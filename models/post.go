package models

type Post struct {
	ID       int      `json:"_id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Images   []string `json:"images"`
	AuthorId int      `json:"author_id"`
}
