package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func RenderAbout(w http.ResponseWriter, r *http.Request) {
	pages.AboutPage().Render(r.Context(), w)
}
