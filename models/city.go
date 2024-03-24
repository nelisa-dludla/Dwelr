package models

type City struct {
	TableName string `gorm:"tableName:cities"`
	ID uint `gorm:"primaryKey"`
	CityName string `gorm:"size:50;not null"`
	ProvinceID uint `gorm:"not null"`
	Province Province
	Suburbs []Suburb `gorm:"foreignKey:CityID"`
}
