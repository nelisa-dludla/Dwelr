package db

import (
	"dwelr/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) (*gorm.DB, error) {
	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Check if tables already exist
	var count int64
	db.Raw("SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public'").Scan(&count)
	tablesExist := count > 0
	if !tablesExist {
		// Migrate models
		db.AutoMigrate(&models.User{}, &models.UserType{}, &models.Region{}, &models.City{}, &models.Area{}, &models.Suburb{}, &models.Listing{}, &models.ListingType{}, &models.Amenity{}, &models.ListingAmenity{})
	}
	// Assign db to package level DB variable
	DB = db
	// Return db instance
	return db, nil
}
