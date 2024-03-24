package models

type Suburb struct {
	TableName string `gorm:"tableName:suburbs"`
	ID uint `gorm:"primaryKey"`
	SuburbName string `gorm:"size:50;not null"`
	CityID uint `gorm:"not null"`
	City City
}
