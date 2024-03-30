package handlers

import (
	"dwelr/models"
	"dwelr/setup"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func RetrieveLocation(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("location")

	type SuburbResult struct {
		SuburbName string `gorm:"suburb_name"`
		CityName string `gorm:"city_name"`
		ProvinceName string `gorm:"province_name"`
	}

	var results []SuburbResult
	setup.DB.Table("suburbs").
		Select("suburbs.suburb_name, cities.city_name, provinces.province_name").
        Joins("JOIN cities ON suburbs.city_id = cities.id").
        Joins("JOIN provinces ON cities.province_id = provinces.id").
        Where("suburbs.suburb_name ILIKE ?", "%"+query+"%").
        Find(&results)
	
	if len(results) < 1 {
		w.Write([]byte("Suburb not found."))
		return
	}
	var resultStr []string
	for _, result := range results {
		locationStr := fmt.Sprintf("%s, %s, %s", result.SuburbName, result.CityName, result.ProvinceName)
		item := fmt.Sprintf("<li onclick='selectInput(this)' class='p-1 w-full hover:cursor-pointer hover:bg-orange-300 hover:rounded hover:font-semibold location_option' data-value='%s'>%s</li>\n", locationStr, locationStr)
		resultStr = append(resultStr, item)
	}
	w.Header().Set("Content-Type", "text/html")
	htmlContent := strings.Join(resultStr, "")
	w.Write([]byte(htmlContent))
}

func RetrieveCitySuburb(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("search")

	// Slice to store location names
	var locations []string
	// Retrieve suburbs
	var provinces []models.Province
	setup.DB.Where("province_name ILIKE ?", "%"+query+"%").Find(&provinces)
	for _, province := range provinces {
		locations = append(locations, province.ProvinceName)
	}
	// Retrieve cities
	var cities []models.City
	setup.DB.Where("city_name ILIKE ?", "%"+query+"%").Find(&cities)
	for _, city := range cities {
		locations = append(locations, city.CityName)
	}
	// Retrieve suburbs
	var suburbs []models.Suburb
	setup.DB.Where("suburb_name ILIKE ?", "%"+query+"%").Find(&suburbs)
	for _, suburb := range suburbs {
		locations = append(locations, suburb.SuburbName)
	}
	sort.Strings(locations)

	if len(locations) < 1 {
		w.Write([]byte("Location not found."))
		return
	}
	var resultStr []string
	for _, location := range locations {
		item := fmt.Sprintf("<li onclick='selectInput(this)' class='p-1 w-full hover:cursor-pointer hover:bg-orange-300 hover:rounded hover:font-semibold location_option' data-value='%s'>%s</li>\n", location, location)
		resultStr = append(resultStr, item)
	}
	w.Header().Set("Content-Type", "text/html")
	htmlContent := strings.Join(resultStr, "")
	w.Write([]byte(htmlContent))
}
