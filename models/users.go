package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"model"`
	UserName   string  `gorm:"column:username;primary_key"`
	Password   string  `gorm:"column:password"`
	Role       string  `gorm:"column:role"`
	CustomerId uint    `gorm:"column:customer_id"`
	CreatedOn  []uint8 `gorm:"column:created_on"`
}
