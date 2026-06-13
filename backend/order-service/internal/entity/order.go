package entity

import "time"

type Order struct {
	OrderID int

	SenderName string
	ReceiverName string

	WeightKg float64

	ShippingCost float64

	TotalPrice float64

	Status string

	CreatedAt time.Time
}