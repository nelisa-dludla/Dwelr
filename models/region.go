package models

type Region struct {
	TableName string `gorm:"tableName:regions"`
	ID uint `gorm:"primaryKey"`
	RegionName string `gorm:"size:50;not null"`
}
