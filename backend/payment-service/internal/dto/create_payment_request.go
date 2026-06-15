package dto

type CreatePaymentRequest struct {
	OrderID int `json:"order_id"`
	PaymentMethod string `json:"payment_method"`

	Total float64 `json:"total"`
}