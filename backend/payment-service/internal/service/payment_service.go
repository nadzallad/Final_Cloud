package service

import (
	"time"

	"payment-service/internal/dto"
	"payment-service/internal/entity"
	"payment-service/internal/rabbitmq"
	"payment-service/internal/repository"

	"github.com/google/uuid"
)

type PaymentService struct {
	Repo      *repository.PaymentRepository
	Publisher *rabbitmq.Publisher
}

func (s *PaymentService) CreatePayment(
	req dto.CreatePaymentRequest,
) error {

	payment := entity.Payment{
		PaymentID:     uuid.New(),
		OrderID:       req.OrderID,
		PaymentMethod: req.Method,
		Total:         req.Total,
		Status:        "PENDING",
		CreatedAt:     time.Now(),
	}

	err := s.Repo.Create(payment)

	if err != nil {
		return err
	}

	event := map[string]interface{}{
		"event": "payment.created",
		"order_id": req.OrderID,
	}

	s.Publisher.Publish(
		"payment.created",
		event,
	)

	return nil
}