package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func ViewListing(w http.ResponseWriter, r *http.Request) {
	pages.ViewListingPage().Render(r.Context(), w)
}
