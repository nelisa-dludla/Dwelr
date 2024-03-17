package models

type Suburb struct {
	TableName string `gorm:"tableName:suburbs"`
	ID uint `gorm:"primaryKey"`
	SuburbName string `gorm:"size:50;not null"`
	AreaID uint `gorm:"not null"`
	Area Area `gorm:"foreignKey:AreaID"`
}
