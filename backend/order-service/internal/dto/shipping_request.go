package dto

type ShippingRequest struct {
	OriginCity string `json:"origin_city"`
	DestinationCity string `json:"destination_city"`

	WeightKg float64 `json:"weight_kg"`

	ServiceType string `json:"service_type"`
}