package handlers

import (
	"dwelr/models"
	"dwelr/setup"
	//"dwelr/views/pages"
	"fmt"
	"net/http"
)

func SearchLocation(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("search")

	var listings []models.Listing

	result := setup.DB.Where("location ILIKE ?", "%"+query+"%").Find(&listings)
	if result.Error != nil {
		fmt.Println("No listings found.")
		return
	}
	fmt.Println("Listings have been found:", listings)

	// Shorten description to 140 characters
	for _, listing := range listings {
		listing.Description = shortenDescription(listing.Description, 140)
	}
	// Render results
	//pages.Results(query, listings).Render(r.Context(), w)

}

func shortenDescription(description string, maxLength int) string {
	if len(description) <= maxLength {
		return description
	}
	return description[:maxLength] + "..."
}
