package entity

import "time"

type Delivery struct {
	DeliveryID      int        `gorm:"column:delivery_id;primaryKey;autoIncrement" json:"delivery_id"`
	TrackingID      string     `gorm:"column:tracking_id;not null" json:"tracking_id"`
	NoResi          string     `gorm:"column:no_resi" json:"no_resi"`
	DeliveryAddress string     `gorm:"column:delivery_address;not null" json:"delivery_address"`
	CourierName     string     `gorm:"column:courier_name" json:"courier_name"`
	CourierPhone    string     `gorm:"column:courier_phone" json:"courier_phone"`
	Status          string     `gorm:"column:status;not null;default:OUT_FOR_DELIVERY" json:"status"`
	DeliveredAt     *time.Time `gorm:"column:delivered_at" json:"delivered_at,omitempty"`
	CreatedAt       time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Delivery) TableName() string {
	return "deliveries"
}