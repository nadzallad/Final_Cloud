package dto

type PickupResponse struct {
	PickupID       int     `json:"pickup_id"`
	UserID         int     `json:"user_id"`
	TrackingNumber string  `json:"tracking_number"`
	PaymentStatus  string  `json:"payment_status"`
	WeightKg       float64 `json:"weight_kg"`
	Status         string  `json:"status"`
}
