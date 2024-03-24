package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	pages.AboutPage().Render(r.Context(), w)
}
