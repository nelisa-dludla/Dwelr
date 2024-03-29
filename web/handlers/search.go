package handlers

import (
	"dwelr/helpers"
	"dwelr/models"
	"dwelr/setup"
	"dwelr/views/pages"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func SearchLocation(w http.ResponseWriter, r *http.Request) {
	// Get search query
	query := r.FormValue("search")
	urlSearch := chi.URLParam(r, "search")
	if query == "" && urlSearch != "" {
		query = urlSearch
	}
	// Get page number
	perPage := 1
	page := 1
	pageStr := chi.URLParam(r, "page")
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	pagination := helpers.PaginationData(page, query, perPage, models.Listing{}, "/search")

	var listings []models.Listing

	result := setup.DB.Limit(perPage).Offset(pagination.Offset).Where("location ILIKE ?", "%"+query+"%").Find(&listings)																	
	if result.Error != nil {
		fmt.Println("No listings found.")
		return
	}

	// Render results
	isLoggedIn, user := LoggedIn(r)
	pages.Results(isLoggedIn, user, query, listings, pagination).Render(r.Context(), w)
}
