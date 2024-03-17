package models

type ListingType struct {
	TableName string `gorm:"tableName:listing_types"`
	ID uint `gorm:"primaryKey"`
	TypeName string `gorm:"size:50;not null"`
}
