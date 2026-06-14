package service

import (
	"fmt"
	"math"
	"order-service/internal/dto"
	"order-service/internal/entity"
	"order-service/internal/repository"
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
	originCity, err := s.CityRepo.GetByID(req.OriginCityID)
	if err != nil {
		return nil, fmt.Errorf("origin city not found: %v", err)
	}

	destCity, err := s.CityRepo.GetByID(req.DestinationCityID)
	if err != nil {
		return nil, fmt.Errorf("destination city not found: %v", err)
	}

	// geocode nama kota → dapat koordinat
	originLat, originLon, err := GetCoordinate(originCity.Name)
	if err != nil {
		return nil, fmt.Errorf("geocode origin failed: %v", err)
	}

	destLat, destLon, err := GetCoordinate(destCity.Name)
	if err != nil {
		return nil, fmt.Errorf("geocode destination failed: %v", err)
	}

	// hitung jarak via OSRM
	distanceKm, err := GetDistance(originLon, originLat, destLon, destLat)
	if err != nil {
		return nil, fmt.Errorf("get distance failed: %v", err)
	}

	shippingCost := calculateShippingCost(req.WeightKg, distanceKm, req.ServiceType)
	totalPrice := shippingCost

	order := &entity.Order{
		UserID:            req.UserID,
		SenderName:        req.SenderName,
		SenderPhone:       req.SenderPhone,
		SenderAddress:     req.SenderAddress,
		ReceiverName:      req.ReceiverName,
		ReceiverPhone:     req.ReceiverPhone,
		ReceiverAddress:   req.ReceiverAddress,
		ItemName:          req.ItemName,
		ItemType:          req.ItemType,
		WeightKg:          req.WeightKg,
		DistanceKm:        math.Round(distanceKm*100) / 100,
		OriginCityID:      req.OriginCityID,
		DestinationCityID: req.DestinationCityID,
		ServiceType:       req.ServiceType,
		BasePrice:         0,
		ShippingCost:      shippingCost,
		TotalPrice:        totalPrice,
		Status:            "WAITING_PAYMENT",
	}

	err = s.OrderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetAllOrders() ([]entity.Order, error) {
	return s.OrderRepo.FindAll()
}

func (s *OrderService) MarkAsPaid(orderID int) error {
	return s.OrderRepo.UpdateStatus(orderID, "PAID")
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
	default: // EZ
		tarifPerKm = 100
	}

	ongkir := roundedWeight * tarifPerKm * distanceKm
	ongkir = math.Round(ongkir)

	if ongkir < 15000 {
		ongkir = 15000
	}

	return ongkir
}