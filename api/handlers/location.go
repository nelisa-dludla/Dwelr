package handlers

import (
	"dwelr/setup"
	"fmt"
	"net/http"
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
		item := fmt.Sprintf("<li onclick='selectInput(this)' class='p-1 w-full hover:cursor-pointer hover:bg-orange-300 location_option' data-value='%s'>%s</li>\n", locationStr, locationStr)
		resultStr = append(resultStr, item)
	}
	w.Header().Set("Content-Type", "text/html")
	htmlContent := strings.Join(resultStr, "")
	w.Write([]byte(htmlContent))
}
