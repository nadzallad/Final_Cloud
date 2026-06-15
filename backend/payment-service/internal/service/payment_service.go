package service

import (
	"bytes"
	"fmt"
	"net/http"
	"payment-service/internal/dto"
	"payment-service/internal/entity"
	"payment-service/internal/rabbitmq"
	"payment-service/internal/repository"
	"strconv"
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

	snapResp, err :=
		CreateSnapTransaction(
			req.OrderID,
			req.Total,
		)

	if err != nil {
		return nil, err
	}

	payment := entity.Payment{
		OrderID: req.OrderID,

		PaymentMethod: req.PaymentMethod,

		Total: req.Total,

		Status: "PENDING",
	}

	err = s.repo.Create(&payment)

	if err != nil {
		return nil, err
	}

	payment.PaymentURL =
		snapResp.RedirectURL

	return &payment, nil
}

func (s *PaymentService) MarkAsPaid(
	orderID string,
) (*entity.Payment, error) {

	id, err := strconv.Atoi(orderID)

	if err != nil {
		return nil, err
	}

	payment, err :=
		s.repo.FindByOrderID(id)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	payment.Status = "PAID"

	err = s.repo.Update(payment)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(
		"http://localhost:8081/api/orders/%s/confirm-payment",
		orderID,
	)

	_, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer([]byte(`{}`)),
	)

	if err != nil {
		return nil, err
	}

	return payment, nil
}
