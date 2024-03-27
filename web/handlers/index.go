package handlers

import (
	"dwelr/models"
	"dwelr/views/pages"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	isLoggedIn, user := LoggedIn(r)
	pages.IndexPage(isLoggedIn, user).Render(r.Context(), w)
}

func LoggedIn(r *http.Request) (bool, models.User) {

	user := r.Context().Value("user")
	if user != nil {
		return true, user.(models.User)
	} else {
		user := models.User{}
		return false, user
	}
}
