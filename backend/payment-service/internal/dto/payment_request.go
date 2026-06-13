package dto

type CreatePaymentRequest struct {
	OrderID string `json:"order_id"`
	Method  string `json:"payment_method"`
	Total   float64 `json:"total"`
}