package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gowaliullah/blog-rest-api/models"
	"github.com/gowaliullah/blog-rest-api/util"
)

var cdb *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	// Decode the JSON payload from the request body into the newUser struct.
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		util.SendError(w, http.StatusBadRequest, "Please provide valid JSON data.")
		return
	}

	// Basic validation to ensure required fields are present.
	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" {
		util.SendError(w, http.StatusBadRequest, "Username, email, and password are required.")
		return
	}

	// SQL query to insert a new user and return the auto-generated ID and timestamps.
	query := `INSERT INTO users (username, email, password, user_role, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING user_id, created_at, updated_at`

	// Execute the query and scan the returned values into the newUser struct.
	// We use QueryRow for a single row result.
	err := cdb.QueryRow(query,
		newUser.Username,
		newUser.Email,
		newUser.Password,
		newUser.UserRole,
		time.Now(),
		time.Now()).Scan(&newUser.UserID, &newUser.CreatedAt, &newUser.UpdatedAt)

	if err != nil {
		// Handle potential errors from the database operation.
		util.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Could not create user: %v", err))
		return
	}

	// For security, remove the password from the response.
	newUser.Password = ""

	// Send a successful JSON response with the new user data.
	util.SendData(w, newUser, http.StatusCreated)
}
