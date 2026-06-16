package entity

import "time"

type Shipment struct {
	ShipmentID      int       `gorm:"column:shipment_id;primaryKey;autoIncrement" json:"shipment_id"`
	TrackingID      string    `gorm:"column:tracking_id" json:"tracking_id"`
	NoResi          string    `gorm:"column:no_resi;unique;not null" json:"no_resi"`
	OriginCity      string    `gorm:"column:origin_city;not null" json:"origin_city"`
	DestinationCity string    `gorm:"column:destination_city;not null" json:"destination_city"`
	CurrentLocation string    `gorm:"column:current_location" json:"current_location"`
	Status          string    `gorm:"column:status;not null;default:IN_TRANSIT" json:"status"`
	ETA             *time.Time `gorm:"column:eta" json:"eta,omitempty"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Shipment) TableName() string {
	return "shipments"
}