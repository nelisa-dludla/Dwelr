package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func AddListing(w http.ResponseWriter, r *http.Request) {
	pages.AddListingPage().Render(r.Context(), w)
}
