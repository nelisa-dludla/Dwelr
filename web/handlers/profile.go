package handlers

import (
	"dwelr/models"
	"dwelr/setup"
	"dwelr/views/pages"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func RenderProfile(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	isLoggedIn, _ := LoggedIn(r)

	var user models.User
	result := setup.DB.First(&user, userId)
	if result.Error != nil {
		pages.NotFound(isLoggedIn, user).Render(r.Context(), w)
		return
	}
	pages.ProfilePage(isLoggedIn, user).Render(r.Context(), w)
}
