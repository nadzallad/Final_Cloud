package dto

type CreateDeliveryRequest struct {
	TrackingID      string `json:"tracking_id" binding:"required"`
	NoResi          string `json:"no_resi"`
	DeliveryAddress string `json:"delivery_address" binding:"required"`
	CourierName     string `json:"courier_name"`
	CourierPhone    string `json:"courier_phone"`
}

type UpdateDeliveryStatusRequest struct {
	Status string `json:"status" binding:"required"`
}