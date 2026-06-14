package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/internal/entity"
	"order-service/internal/dto"
	"order-service/internal/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func NewOrderService(
	repo *repository.OrderRepository,
) *OrderService {

	return &OrderService{
		Repo: repo,
	}
}

func (s *OrderService) MarkAsPaid(
	orderID string,
) error {

	return s.Repo.UpdateStatus(
		orderID,
		"PAID",
	)
}

func (s *OrderService) CreateOrder(
	req dto.CreateOrderRequest,
) (map[string]interface{}, error) {

	// simulasi perhitungan ongkir
	totalPrice := req.WeightKg * 10000

	// sementara hardcode
	orderID := "ORD001"

	paymentReq := map[string]interface{}{
		"order_id": orderID,
		"total":    totalPrice,
		"payment_method": req.PaymentMethod,
	}

	jsonData, err := json.Marshal(paymentReq)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		"http://localhost:8082/payments",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// cek response payment
	if resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf(
			"payment service returned status %d",
			resp.StatusCode,
		)
	}

	return map[string]interface{}{
		"order_id": orderID,
		"status":   "PENDING_PAYMENT",
		"total":    totalPrice,
	}, nil
}

func (s *OrderService) GetOrders() (
	[]entity.Order,
	error,
) {

	return s.Repo.FindAll()
}