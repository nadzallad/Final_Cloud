package entity

import "time"

type Pickup struct {
	PickupID       int       `gorm:"column:pickup_id;primaryKey;autoIncrement" json:"pickup_id"`
	UserID         int       `gorm:"column:user_id" json:"user_id"`
	TrackingNumber string    `gorm:"column:tracking_number;unique" json:"tracking_number"`
	PaymentStatus  string    `gorm:"column:payment_status" json:"payment_status"`
	WeightKg       float64   `gorm:"column:weight_kg" json:"weight_kg"`
	Status         string    `gorm:"column:status;default:WAITING_PICKUP" json:"status"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Pickup) TableName() string {
	return "pickups"
}
