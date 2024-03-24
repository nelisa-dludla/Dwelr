package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pages.IndexPage().Render(r.Context(), w)
}
