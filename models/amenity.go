package models

type Amenity struct {
	TableName string `gorm:"tableName:amenities"`
	ID uint `gorm:"primaryKey"`
	AmenityName string `gorm:"size:50;not null"`
}
