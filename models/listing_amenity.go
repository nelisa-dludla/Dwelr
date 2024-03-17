package models

type ListingAmenity struct {
	TableName string `gorm:"tableName:listing_amenities"`
	ListingID uint `gorm:"primaryKey"`
	AmenityID uint `gorm:"primaryKey"`
}
