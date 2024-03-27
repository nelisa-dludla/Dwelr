package handlers

import (
	"dwelr/models"
	"dwelr/setup"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SearchLocation(w http.ResponseWriter, r *http.Request) {
	search := chi.URLParam(r, "search")

	var listings []models.Listing

	result := setup.DB.Where("location ILIKE ?", "%"+search+"%").Find(&listings)
	if result.Error != nil {
		fmt.Println("No listings found.")
		return
	}
	fmt.Println("Listings have been found:", listings)

	var listingCardTemplate = template.Must(template.New("listingCard").Parse(`
		<a href="/view-listing/{{ .ID }}">
			<div class="flex h-72 w-full mt-8 max-w-6xl rounded shadow-xl hover:cursor-pointer hover:shadow-lime-800">
    			<div class="w-2/5">
      				<figure class="bg-black flex justify-center items-center h-full w-full">
        				<img class="h-full" src="/views/images/Placeholder.svg"/>
      				</figure>
    			</div>
    			<div class="flex flex-col justify-center w-3/5 p-8">
      				<p class="font-bold text-2xl">R {{.RentalFee}}</p>
      				<p class="my-4">{{ .NumberOfBedrooms }} Bedroom Apartment</p>
      				<p class="font-bold">{{ .Location }}</p>
      				<p class="my-4">{{ .Description }}</p>
      				<div class="flex font-bold my-4">
      				  <p>{{ .NumberOfBedrooms }} Bedroom</p>
      				  <p class="ml-1">{{ .NumberOfBathrooms }} Bathroom</p>
      				</div>
  				</div>
  			</div>
		</a>
	`))

	var noResultsTemplate = template.Must(template.New("NoResults").Parse(`
		<div>
			<h2 class="text-xl">No results</h2>
		</div>
	`))

	if len(listings) < 1 {
		err := noResultsTemplate.Execute(w, listings)
		if err != nil {
			log.Fatal("Failed to render noResultsTemplate:", err)
			return
		}
	}
	for _, listing := range listings {
		// Shorten description
		listing.Description = shortenDescription(listing.Description, 140)
		err := listingCardTemplate.Execute(w, listing)
		if err != nil {
			log.Fatal("Failed to render listingCardTemplate:", err)
			return
		}
	}
}

func shortenDescription(description string, maxLength int) string {
	if len(description) <= maxLength {
		return description
	}
	return description[:maxLength] + "..."
}
