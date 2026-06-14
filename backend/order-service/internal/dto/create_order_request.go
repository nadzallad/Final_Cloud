package dto

type CreateOrderRequest struct {
	UserID int `json:"user_id"`

	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
	SenderAddress string `json:"sender_address"`

	ReceiverName string `json:"receiver_name"`
	ReceiverPhone string `json:"receiver_phone"`
	ReceiverAddress string `json:"receiver_address"`

	ItemName string `json:"item_name"`
	ItemType string `json:"item_type"`

	WeightKg float64 `json:"weight_kg"`

	OriginCity string `json:"origin_city"`
	DestinationCity string `json:"destination_city"`

	ServiceType string `json:"service_type"`

	PaymentMethod string `json:"payment_method"`
}