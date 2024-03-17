package handlers

import (
	"dwelr/db"
	"dwelr/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)
// Retrieves a users data from the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(user)
}
// Retrieves all of a users listings from the database
func GetUserListings(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
	}

	var listings []models.Listing
	result := db.DB.Where("id = ?", id).Find(&listings)
	if result.Error != nil {
		http.Error(w, "Could not find any listings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(&listings)
}
