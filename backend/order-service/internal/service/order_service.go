package service

import (
	"order-service/internal/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func (s *OrderService) MarkAsPaid(
	orderID int,
) error {

	return s.Repo.UpdateStatus(
		orderID,
		"PAID",
	)
}