package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func RenderViewListing(w http.ResponseWriter, r *http.Request) {
	pages.ViewListingPage().Render(r.Context(), w)
}
