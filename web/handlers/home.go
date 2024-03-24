package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	pages.HomePage().Render(r.Context(), w)
}
