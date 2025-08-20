package models

type Category struct {
	ID   int    `json:"_id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
