package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	isLoggedIn, user := LoggedIn(r)
	pages.AboutPage(isLoggedIn, user).Render(r.Context(), w)
}
