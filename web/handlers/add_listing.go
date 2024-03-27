package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func AddListing(w http.ResponseWriter, r *http.Request) {
	isLoggedIn, user := LoggedIn(r)
	pages.AddListingPage(isLoggedIn, user).Render(r.Context(), w)
}
