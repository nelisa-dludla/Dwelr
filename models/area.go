package models

type Area struct {
	TableName string `gorm:"tableName:areas"`
	ID uint `gorm:"primaryKey"`
	AreaName string `gorm:"size:50;not null"`
	CityID uint `gorm:"not null"`
	City City `gorm:"foreignKey:CityID"`
}
