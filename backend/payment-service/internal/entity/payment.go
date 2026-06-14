package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	PaymentID uuid.UUID `gorm:"column:payment_id;primaryKey"`

	OrderID string `gorm:"column:order_id"`

	PaymentMethod string `gorm:"column:payment_method"`

	Total float64 `gorm:"column:total"`

	Discount float64 `gorm:"column:discount"`

	AdminFee float64 `gorm:"column:admin_fee"`

	Status string `gorm:"column:status"`

	CreatedAt time.Time `gorm:"column:created_at"`

	PaidAt *time.Time `gorm:"column:paid_at"`

	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}