package dto

type CreatePickupRequest struct {
	UserID         int     `json:"user_id"`
	TrackingNumber string  `json:"tracking_number"`
	PaymentStatus  string  `json:"payment_status"`
	WeightKg       float64 `json:"weight_kg"`
}
