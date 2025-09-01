package users

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gowaliullah/blog-rest-api/models"
	"github.com/gowaliullah/blog-rest-api/util"
)

var db *sql.DB

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	query := `SELECT user_id, username, email, password, user_role, created_at, updated_at FROM users`

	rows, err := db.Query(query)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, fmt.Sprintf("could not get all users: %v", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.UserRole, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, fmt.Sprintf("could not scan user row: %v", err))
			return
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		util.SendError(w, http.StatusInternalServerError, fmt.Sprintf("error in rows iteration: %v", err))
		return
	}

	util.SendData(w, users, http.StatusOK)
}
