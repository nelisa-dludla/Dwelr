package models

import "time"

type User struct {
	TableName string `gorm:"tableName:users"`
	ID uint `gorm:"primaryKey"`
	FirstName string `gorm:"size:50;not null"`
	LastName string `gorm:"size:50;not null"`
	Email string `gorm:"size:100;unique;not null"`
	Password string `gorm:"size:100;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Listings []Listing `gorm:"foreignKey:UserID"`
}
