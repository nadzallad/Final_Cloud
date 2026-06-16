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

	fmt.Println("===== CREATE PAYMENT =====")
	fmt.Println("OrderID:", req.OrderID)
	fmt.Println("Method :", req.PaymentMethod)
	fmt.Println("Total  :", req.Total)
	fmt.Println("REAL ORDER ID:", req.OrderID)


	payment := entity.Payment{
		OrderID: req.OrderID,
		PaymentMethod: req.PaymentMethod,
		Total: req.Total,
		Status: "PENDING",
	}

	fmt.Println("SAVE DATABASE...")

	if err := s.repo.Create(&payment); err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}

	fmt.Println("PAYMENT SAVED")
	fmt.Println("PaymentID:", payment.PaymentID)

	orderID := strconv.Itoa(
		req.OrderID,
	)

	fmt.Println("MIDTRANS ORDER ID:", orderID)

	fmt.Println("CALL MIDTRANS...")

	snapResp, err := CreateSnapTransaction(
		orderID,
		req.Total,
	)

	if err != nil {

		fmt.Println("MIDTRANS ERROR:")
		fmt.Println(err)

		return nil, err
	}

	fmt.Println("MIDTRANS SUCCESS")
	fmt.Println("TOKEN:", snapResp.Token)
	fmt.Println("URL:", snapResp.RedirectURL)

	payment.PaymentURL =
		snapResp.RedirectURL

	fmt.Println("RETURN RESPONSE")

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
