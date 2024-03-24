package setup

import (
	"dwelr/models"

	"gorm.io/gorm"
)

func SyncDb(DB *gorm.DB) {
	// Check if tables already exist
	var count int64
	DB.Raw("SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public'").Scan(&count)
	tablesExist := count > 0
	if !tablesExist {
		// Migrate models
		DB.AutoMigrate(&models.User{}, &models.Province{}, &models.City{}, &models.Suburb{}, &models.Listing{})
	}
}
