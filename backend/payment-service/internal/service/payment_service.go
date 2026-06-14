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
	repo      *repository.PaymentRepository
	publisher *rabbitmq.Publisher
}

func NewPaymentService(
	repo *repository.PaymentRepository,
	publisher *rabbitmq.Publisher,
) *PaymentService {

	return &PaymentService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *PaymentService) CreatePayment(
	req dto.CreatePaymentRequest,
) (*entity.Payment, error) {

	payment := entity.Payment{
		PaymentID: uuid.New(),

		OrderID: req.OrderID,

		Total: req.Total,

		Discount: 0,

		AdminFee: 0,

		Status: "UNPAID",

		PaymentMethod: req.PaymentMethod,
	}

	err := s.repo.Create(&payment)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (s *PaymentService) MarkAsPaid(
	paymentID string,
) (*entity.Payment, error) {

	payment, err :=
		s.repo.FindByID(paymentID)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	payment.Status = "PAID"
	payment.PaidAt = &now

	err = s.repo.Update(payment)

	if err != nil {
		return nil, err
	}

	if s.publisher != nil {

		err = s.publisher.Publish(
			"payment.success",
			map[string]interface{}{
				"order_id": payment.OrderID,
				"status":   "PAID",
			},
		)

		if err != nil {
			return nil, err
		}
	}

	return payment, nil
}