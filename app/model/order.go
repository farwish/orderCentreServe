package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)
// Note: We relay on golang-migrate/migrate other than gorm tags AutoMigrate

// Models define: https://gorm.io/zh_CN/docs/models.html
type Order struct {
	OrderId 	string `gorm:"primaryKey"`

	AccountId 	uint `gorm:"index"`
	ProductId 	uint `gorm:"index"`
	TradeNo 	string
	TotalFee 	uint
	Status 		uint
	PaymentType uint
	PaymentStatus uint

	PaymentAt 	sql.NullTime

	Source 		string
	ProductName string

	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}