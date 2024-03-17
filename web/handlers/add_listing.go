package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func RenderAddListing(w http.ResponseWriter, r *http.Request) {
	pages.AddListingPage().Render(r.Context(), w)
}
