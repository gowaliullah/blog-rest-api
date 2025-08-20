package models

import (
	"time"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleAuthor UserRole = "author"
	RoleReader UserRole = "reader"
)

type User struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UserRole  UserRole  `json:"user_role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type User struct {
// 	UserID    int       `db:"user_id" json:"user_id"`
// 	Username  string    `db:"username" json:"username"`
// 	Email     string    `db:"email" json:"email"`
// 	Password  string    `db:"password" json:"password"`
// 	UserRole  UserRole  `db:"user_role" json:"user_role"`
// 	CreatedAt time.Time `db:"created_at" json:"created_at"`
// 	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
// }
