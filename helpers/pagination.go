package helpers

import (
	"dwelr/models"
	"dwelr/setup"
	"math"
)

func PaginationData(page int, query string, perPage int, model interface{}, baseURL string) models.Pagination {
	// Calculate total pages
	var totalRows int64
	setup.DB.Model(model).Where("location ILIKE ?", "%"+query+"%").Count(&totalRows)
	totalPages := math.Ceil(float64(float64(totalRows)/float64(perPage)))

	// Calculate offset
	offset := (page - 1) * perPage

	return models.Pagination{
		NextPage: page + 1,
		PrevPage: page - 1,
		CurrentPage: page,
		TotalPages: int(totalPages),
		TwoBelow: page - 2,
		TwoAfter: page + 2,
		ThreeAfter: page + 3,
		Offset: offset,
		BaseURL: baseURL,
	}
}
