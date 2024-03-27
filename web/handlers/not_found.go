package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func Not_Found(w http.ResponseWriter, r *http.Request) {
	isLoggedIn, user := LoggedIn(r)
	pages.NotFound(isLoggedIn, user).Render(r.Context(), w)
}
