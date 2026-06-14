package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"order-service/internal/dto"
	"order-service/internal/entity"
	"order-service/internal/repository"
	"strconv"
	"time"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
	CityRepo  *repository.CityRepository
}

func NewOrderService(
	orderRepo *repository.OrderRepository,
	cityRepo *repository.CityRepository,
) *OrderService {
	return &OrderService{
		OrderRepo: orderRepo,
		CityRepo:  cityRepo,
	}
}

func (s *OrderService) CreateOrder(req dto.CreateOrderRequest) (*entity.Order, error) {
	originLat, originLon, err := GetCoordinate(req.OriginCity)
	if err != nil {
		return nil, fmt.Errorf("geocode origin failed: %v", err)
	}

	destLat, destLon, err := GetCoordinate(req.DestinationCity)
	if err != nil {
		return nil, fmt.Errorf("geocode destination failed: %v", err)
	}

	distanceKm, err := GetDistance(originLon, originLat, destLon, destLat)
	if err != nil {
		return nil, fmt.Errorf("get distance failed: %v", err)
	}

	shippingCost := calculateShippingCost(req.WeightKg, distanceKm, req.ServiceType)
	totalPrice := shippingCost

	order := &entity.Order{
		UserID:          req.UserID,
		SenderName:      req.SenderName,
		SenderPhone:     req.SenderPhone,
		SenderAddress:   req.SenderAddress,
		ReceiverName:    req.ReceiverName,
		ReceiverPhone:   req.ReceiverPhone,
		ReceiverAddress: req.ReceiverAddress,
		ItemName:        req.ItemName,
		ItemType:        req.ItemType,
		WeightKg:        req.WeightKg,
		DistanceKm:      math.Round(distanceKm*100) / 100,
		OriginCity:      req.OriginCity,
		DestinationCity: req.DestinationCity,
		ServiceType:     req.ServiceType,
		ShippingCost:    shippingCost,
		TotalPrice:      totalPrice,
		Status:          "WAITING_PAYMENT",
	}

	err = s.OrderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	// kirim ke Payment Service
	paymentReq := map[string]interface{}{
		"order_id":       strconv.Itoa(order.OrderID),
		"payment_method": req.PaymentMethod,
		"total":          totalPrice,
	}

	jsonData, _ := json.Marshal(paymentReq)

	resp, err := http.Post(
		"http://localhost:8082/payments",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		fmt.Println("Warning: gagal kirim ke payment service:", err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
			fmt.Println("Warning: payment service returned status:", resp.StatusCode)
		}
	}

	return order, nil
}

func (s *OrderService) GetAllOrders() ([]entity.Order, error) {
	return s.OrderRepo.FindAll()
}

func (s *OrderService) MarkAsPaid(orderID string) error {
	id, err := strconv.Atoi(orderID)
	if err != nil {
		return err
	}

	err = s.OrderRepo.UpdateStatus(id, "PAID")
	if err != nil {
		return err
	}

	// generate resi
	noResi := fmt.Sprintf("LOG-%s-%d", orderID, time.Now().Unix())
	err = s.OrderRepo.CreateResi(id, noResi)
	if err != nil {
		return err
	}

	return nil
}

func calculateShippingCost(weightKg float64, distanceKm float64, serviceType string) float64 {
	roundedWeight := math.Ceil(weightKg)

	var tarifPerKm float64
	switch serviceType {
	case "JSD":
		tarifPerKm = 200
	case "JND":
		tarifPerKm = 150
	case "DOC":
		tarifPerKm = 80
	default:
		tarifPerKm = 100
	}

	ongkir := roundedWeight * tarifPerKm * distanceKm
	ongkir = math.Round(ongkir)

	if ongkir < 15000 {
		ongkir = 15000
	}

	return ongkir
}