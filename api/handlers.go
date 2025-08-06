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
	rows, err := db.Query("SELECT id, name FROM categories")
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
func GetSingleCategory(db *sql.DB, w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category models.Category
	err = db.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Category not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch category", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// creates a new item
func CreateCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var categories models.Category
	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil || categories.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO categories (name) VALUES ($1)", categories.Name)
	if err != nil {
		http.Error(w, "Failed to insert categories item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Category is created")
}

// Updatecategory updates an existing item
func UpdateCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var categories models.Category
	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil || categories.ID == 0 || categories.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("UPDATE categories SET name = $1 WHERE id = $2", categories.Name, categories.ID)
	if err != nil {
		http.Error(w, "Failed to update category name", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No category found to update", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "category updated successfully")
}

// deletes category by ID
func DeleteCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var categoryToDelete models.Category

	// Decode the request body into the Category struct
	if err := json.NewDecoder(r.Body).Decode(&categoryToDelete); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the input ID
	if categoryToDelete.ID == 0 {
		http.Error(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	// Execute the DELETE statement
	result, err := db.Exec("DELETE FROM categories WHERE id = $1", categoryToDelete.ID)
	if err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error checking rows affected", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("No category with ID %d found to delete", categoryToDelete.ID), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Category deleted successfully")
}
