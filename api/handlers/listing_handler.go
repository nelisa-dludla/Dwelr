package handlers

import (
	"dwelr/db"
	"dwelr/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)
// Listing struct to be used when querying the database
var listing models.Listing
// Retrieves listing based on id from the database
func GetListing(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Listing ID", http.StatusBadRequest)
	}

	result := db.DB.First(&listing, id)
	if result.Error != nil {
		http.Error(w, "Listing not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(listing)
}
// Add listing to the database
func AddListing(w http.ResponseWriter, r *http.Request) {
	// location := r.FormValue("location")
	title := r.FormValue("title")
	description := r.FormValue("description")
	rentalFeeStr := r.FormValue("rental_fee")
	rentalFee, _ := strconv.ParseFloat(rentalFeeStr, 64)
	dwellingTypeStr := r.FormValue("dwelling_type")
	dwellingType, _ := strconv.Atoi(dwellingTypeStr)
	// numOfBedroomsStr := r.FormValue("num_of_bedrooms")
	// numOfBedrooms, _ := strconv.Atoi(numOfBedroomsStr)
	// numOfBathroomsStr := r.FormValue("num_of_bathrooms")
	// numOfBathrooms, _ := strconv.Atoi(numOfBathroomsStr)

	listing := models.Listing {
		Title: title,
		Description: description,
		RentalFee: rentalFee,
		ListingTypeID: uint(dwellingType),
	}

	result := db.DB.Create(listing)
	if result.Error != nil {
		http.Error(w, "Failed to add listing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Listing created successfully!")
}
// Struct for updating a listing
type ListingUpdate struct {
	Title string `json:"title"`
	Description string `json:"description"`
	RentalFee float64 `json:"rental_fee"`
}
// Updates listing
func UpdateListing(w http.ResponseWriter, r *http.Request) {
	var updataData ListingUpdate

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updataData)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid listing ID", http.StatusBadRequest)
		return
	}

	result := db.DB.First(&listing, id)
	if result.Error != nil {
		http.Error(w, "Listing not found", http.StatusNotFound)
		return
	}

	if updataData.Title != "" {
		listing.Title = updataData.Title
	}
	if updataData.Description != "" {
		listing.Description = updataData.Description
	}
	if updataData.RentalFee != 0 {
		listing.RentalFee = updataData.RentalFee
	}

	saveResult := db.DB.Save(&listing)
	if saveResult.Error != nil {
		http.Error(w, "Failed to update listing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Listing updated successfully!")
}
// Deletes a listing from the database
func DeleteListing(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid listing ID", http.StatusBadRequest)
		return
	}

	result := db.DB.First(&listing, id)
	if result.Error != nil {
		http.Error(w, "Listing not found", http.StatusNotFound)
		return
	}

	deleteResult := db.DB.Delete(&listing)
	if deleteResult.Error != nil {
		http.Error(w, "Failed to delete listing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Listing deleted successfully!")
}
// Retrieves all listings from the database
func GetListings(w http.ResponseWriter, r *http.Request) {
	var listings []models.Listing
	result := db.DB.Find(&listings)
	if result.Error != nil {
		http.Error(w, "No listings currently available", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(listings)
}
