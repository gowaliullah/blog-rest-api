package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gowaliullah/blog-rest-api/models"
)

// GetAllCategories fetches all categories items from the database
func GetAllCategories(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, item FROM categories")
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var m models.Category
		if err := rows.Scan(&m.ID, &m.Name); err != nil {
			http.Error(w, "Error scanning categories", http.StatusInternalServerError)
			return
		}
		categories = append(categories, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// GetSingle fetches a single item by ID
func GetSingleCategiry(db *sql.DB, w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var categories models.Category
	err = db.QueryRow("SELECT id, item FROM menus WHERE id = $1", id).Scan(&categories.ID, &categories.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "categories item not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch categories item", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// creates a new item
func CreateCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var categories models.Category
	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil || categories.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO categories (item) VALUES ($1)", categories.Name)
	if err != nil {
		http.Error(w, "Failed to insert categories item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Category item is created")
}

// Updatecategory updates an existing item
func UpdateCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var categories models.Category
	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil || categories.ID == 0 || categories.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("UPDATE categories SET item = $1 WHERE id = $2", categories.Name, categories.ID)
	if err != nil {
		http.Error(w, "Failed to update category item", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No menu item found to update", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Menu item updated successfully")
}

// deletes category by ID
func DeleteCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var data struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil || data.ID == 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM categories WHERE id = $1", data.ID)
	if err != nil {
		http.Error(w, "Failed to delete category item", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No category item found to delete", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "category deleted successfully")
}
