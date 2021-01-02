package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// https://gorm.io/zh_CN/docs/models.html

type Order struct {
	OrderId string `gorm:"primaryKey"`
	AccountId uint
	ProductId uint
	TradeNo uint
	TotalFee uint
	Status uint
	PaymentType uint
	PaymentStatus uint
	PaymentAt sql.NullTime
	ProductName string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}