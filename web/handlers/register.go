package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func RenderRegister(w http.ResponseWriter, r *http.Request) {
	pages.RegisterPage().Render(r.Context(), w)
}
