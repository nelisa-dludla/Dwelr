package models

type Listing struct {
	TableName string `gorm:"tableName:listings"`
	ID uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	User User `gorm:"foreignKey:UserID"`
	Location string `gorm:"size:200;not null"`
	Title string `gorm:"size:100;not null"`
	Description string `gorm:"type:text;not null"`
	RentalFee float64 `gorm:"type:decimal(10,2);not null"`
	Type string `gorm:"size:50;not null"`
	NumberOfBedrooms uint `gorm:"not null"`
	NumberOfBathrooms uint `gorm:"not null"`
}
