package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model  `json:"model"`
	AccountId   uint           `gorm:"column:account_id;primaryKey;autoIncrement"`
	CustomerId  uint           `gorm:"column:customer_id"`
	OpeningDate datatypes.Time `gorm:"column:opening_date"`
	AccountType string         `gorm:"column:account_type"`
	Amount      float32        `gorm:"column:amount;type:decimal(10,2)"`
	Status      bool           `gorm:"column:status"`
}
