package models

type User struct {
	ID       int     `json:"_id"`
	Name     string  `json:"title"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Photo    *string `json:"photo,omitempty"`
}
