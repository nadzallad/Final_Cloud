package dto

import "time"

type CreateShipmentRequest struct {
	TrackingID      string     `json:"tracking_id"`
	NoResi          string     `json:"no_resi" binding:"required"`
	OriginCity      string     `json:"origin_city" binding:"required"`
	DestinationCity string     `json:"destination_city" binding:"required"`
	CurrentLocation string     `json:"current_location"`
	ETA             *time.Time `json:"eta,omitempty"`
}

type UpdateShipmentStatusRequest struct {
	Status          string `json:"status" binding:"required"`
	CurrentLocation string `json:"current_location"`
}

type ShipmentResponse struct {
	ShipmentID      int        `json:"shipment_id"`
	TrackingID      string     `json:"tracking_id"`
	NoResi          string     `json:"no_resi"`
	OriginCity      string     `json:"origin_city"`
	DestinationCity string     `json:"destination_city"`
	CurrentLocation string     `json:"current_location"`
	Status          string     `json:"status"`
	ETA             *time.Time `json:"eta,omitempty"`
	CreatedAt       string     `json:"created_at"`
}