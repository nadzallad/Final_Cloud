package entity

import "time"

type Order struct {
	OrderID         int       `gorm:"column:order_id;primaryKey;autoIncrement" json:"order_id"`
	UserID          int       `gorm:"column:user_id" json:"user_id"`
	SenderName      string    `gorm:"column:sender_name" json:"sender_name"`
	SenderPhone     string    `gorm:"column:sender_phone" json:"sender_phone"`
	SenderAddress   string    `gorm:"column:sender_address" json:"sender_address"`
	ReceiverName    string    `gorm:"column:receiver_name" json:"receiver_name"`
	ReceiverPhone   string    `gorm:"column:receiver_phone" json:"receiver_phone"`
	ReceiverAddress string    `gorm:"column:receiver_address" json:"receiver_address"`
	ItemName        string    `gorm:"column:item_name" json:"item_name"`
	ItemType        string    `gorm:"column:item_type" json:"item_type"`
	WeightKg        float64   `gorm:"column:weight_kg" json:"weight_kg"`
	DistanceKm      float64   `gorm:"column:distance_km" json:"distance_km"`
	OriginCity      string    `gorm:"column:origin_city" json:"origin_city"`
	DestinationCity string    `gorm:"column:destination_city" json:"destination_city"`
	ServiceType     string    `gorm:"column:service_type" json:"service_type"`
	ShippingCost    float64   `gorm:"column:shipping_cost" json:"shipping_cost"`
	TotalPrice      float64   `gorm:"column:total_price" json:"total_price"`
	Status          string    `gorm:"column:status;default:WAITING_PAYMENT" json:"status"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Order) TableName() string {
	return "orders"
}