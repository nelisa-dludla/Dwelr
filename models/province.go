package models

type Province struct {
	TableName string `gorm:"tableName:regions"`
	ID uint `gorm:"primaryKey"`
	ProvinceName string `gorm:"size:50;not null"`
	Cities []City `gorm:"foreignKey:ProvinceID"`
}
