package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	pages.LoginPage().Render(r.Context(), w)
}
