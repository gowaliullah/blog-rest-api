package models

type Menu struct {
	ID       int    `json:"_id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Slug     string `json:"slug"`
	ParentID *int   `json:"parent_id,omitempty"`
}
