package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	pages.RegisterPage().Render(r.Context(), w)
}
