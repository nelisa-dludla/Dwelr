package models

type Listing struct {
	TableName string `gorm:"tableName:listings"`
	ID uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	User User `gorm:"foreignKey:UserID"`
	SuburbID uint `gorm:"not null"`
	Suburb Suburb `gorm:"foreignKey:SuburbID"`
	Title string `gorm:"size:100;not null"`
	Description string `gorm:"type:text;not null"`
	RentalFee float64 `gorm:"type:decimal(10,2);not null"`
	ListingTypeID uint `gorm:"not null"`
	ListingType ListingType `gorm:"foreignKey:ListingTypeID"`
}
