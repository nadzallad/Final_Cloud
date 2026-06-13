package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	PaymentID uuid.UUID `json:"payment_id"`

	OrderID string `json:"order_id"`

	PaymentMethod string `json:"payment_method"`

	Total float64 `json:"total"`

	Discount float64 `json:"discount"`

	AdminFee float64 `json:"admin_fee"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`

	PaidAt *time.Time `json:"paid_at"`
}