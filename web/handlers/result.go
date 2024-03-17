package handlers

import (
	"dwelr/views/pages"
	"net/http"
)

func RenderResult(w http.ResponseWriter, r *http.Request) {
	pages.ResultPage().Render(r.Context(), w)
}
