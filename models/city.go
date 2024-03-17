package models

type City struct {
	TableName string `gorm:"tableName:cities"`
	ID uint `gorm:"primaryKey"`
	CityName string `gorm:"size:50;not null"`
	RegionID uint `gorm:"not null"`
	Region Region `gorm:"foreignKey:RegionID"`
}
