package handlers

import (
	"dwelr/models"
	"dwelr/setup"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SearchLocation(w http.ResponseWriter, r *http.Request) {
	search := chi.URLParam(r, "search")

	var suburbs []models.Suburb
	var cities []models.City


	suburbResult := setup.DB.Where("name LIKE ?", "%"+search+"%").Find(&suburbs)
	if suburbResult.Error != nil {
		http.Error(w, "No suburb found", http.StatusNotFound)
	}
	cityResult := setup.DB.Where("name LIKE ?", "%"+search+"%").Find(&cities)
	if cityResult.Error != nil {
		http.Error(w, "No city found", http.StatusNotFound)
	}

	fmt.Println(cityResult)

	json.NewEncoder(w).Encode(cityResult)
}
