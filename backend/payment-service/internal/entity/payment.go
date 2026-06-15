package entity

import (
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

func (p *Payment) BeforeCreate(
    tx *gorm.DB,
) error {

    if p.PaymentID == uuid.Nil {
        p.PaymentID = uuid.New()
    }

    return nil
}
type Payment struct {
	PaymentID uuid.UUID `gorm:"column:payment_id;primaryKey"`

	OrderID int  `gorm:"column:order_id"`

	PaymentMethod string `gorm:"column:payment_method"`

	Total float64 `gorm:"column:total"`

	Discount float64 `gorm:"column:discount"`

	AdminFee float64 `gorm:"column:admin_fee"`

	Status string `gorm:"column:status"`

	PaymentURL string `json:"payment_url" gorm:"-"`

	CreatedAt time.Time `gorm:"column:created_at"`

	PaidAt *time.Time `gorm:"column:paid_at"`

	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}