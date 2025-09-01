package users

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/models"
)

var db *sql.DB

func GetAllUsers(w http.ResponseWriter, r *http.Request) ([]models.User, error) {
	var users []models.User
	query := `SELECT user_id, username, email, password, user_role, created_at, updated_at FROM users`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not get all users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.UserRole, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("could not scan user row: %v", err)
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows iteration: %v", err)
	}

	w.WriteHeader(http.StatusNotFound)
	return users, nil
}
