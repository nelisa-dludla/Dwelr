package models

type UserType struct {
	TableName string `gorm:"tableName:user_types"`
	ID uint `gorm:"primaryKey"`
	TypeName string `gorm:"size:50;not null"`
}
