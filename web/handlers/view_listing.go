package handlers

import (
	"dwelr/models"
	"dwelr/setup"
	"dwelr/views/pages"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ViewListing(w http.ResponseWriter, r *http.Request) {
	listingIdStr := chi.URLParam(r, "id")
	listingId, err := strconv.Atoi(listingIdStr)
	if err != nil {
		http.Error(w, "Invalid Listing ID", http.StatusBadRequest)
	}

	var listing models.Listing
	// Retrieve listing based on id from the database
	result := setup.DB.First(&listing, listingId)
	if result.Error != nil {
		http.Error(w, "Listing not found", http.StatusNotFound)
		return
	}

	isLoggedIn, user := LoggedIn(r)
	// Ignore text-editor error message
	pages.ViewListingPage(listing, isLoggedIn, user).Render(r.Context(), w)
}
