package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func Not_Found(w http.ResponseWriter, r *http.Request) {
	pages.NotFound().Render(r.Context(), w)
}
