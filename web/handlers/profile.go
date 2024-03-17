package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func RenderProfile(w http.ResponseWriter, r *http.Request) {
	pages.ProfilePage().Render(r.Context(), w)
}
