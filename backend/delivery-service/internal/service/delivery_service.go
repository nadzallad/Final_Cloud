package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"delivery-service/internal/dto"
	"delivery-service/internal/entity"
	"delivery-service/internal/rabbitmq"
	"delivery-service/internal/repository"
)

type DeliveryService struct {
	repo      *repository.DeliveryRepository
	publisher *rabbitmq.Publisher
}

func NewDeliveryService(repo *repository.DeliveryRepository, publisher *rabbitmq.Publisher) *DeliveryService {
	return &DeliveryService{
		repo:      repo,
		publisher: publisher,
	}
}

func orderServiceURL() string {
	url := os.Getenv("ORDER_SERVICE_URL")
	if url == "" {
		url = "http://order-service:8081"
	}
	return url
}

type receiverDetail struct {
	ReceiverAddress string `json:"receiver_address"`
	ReceiverName    string `json:"receiver_name"`
	ReceiverPhone   string `json:"receiver_phone"`
}

func fetchReceiverDetail(orderID string) (*receiverDetail, error) {
	url := fmt.Sprintf("%s/api/orders/%s", orderServiceURL(), orderID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("order-service returned status %d", resp.StatusCode)
	}

	var data struct {
		ReceiverAddress string `json:"receiver_address"`
		ReceiverName    string `json:"receiver_name"`
		ReceiverPhone   string `json:"receiver_phone"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &receiverDetail{
		ReceiverAddress: data.ReceiverAddress,
		ReceiverName:    data.ReceiverName,
		ReceiverPhone:   data.ReceiverPhone,
	}, nil
}

// CreateDeliveryFromShipment dipanggil dari consumer shipment.delivered
func (s *DeliveryService) CreateDeliveryFromShipment(noResi string, trackingID string) error {
	// Cek duplikat
	existing, _ := s.repo.FindByNoResi(noResi)
	if existing != nil {
		return nil
	}

	address := "Unknown"
	courierName := "Auto-assigned"
	courierPhone := ""

	if trackingID != "" {
		if detail, err := fetchReceiverDetail(trackingID); err == nil {
			if detail.ReceiverAddress != "" {
				address = detail.ReceiverAddress
			}
		}
	}

	delivery := &entity.Delivery{
		TrackingID:      trackingID,
		NoResi:          noResi,
		DeliveryAddress: address,
		CourierName:     courierName,
		CourierPhone:    courierPhone,
		Status:          "OUT_FOR_DELIVERY",
	}

	return s.repo.Create(delivery)
}

// CreateDelivery - manual via REST
func (s *DeliveryService) CreateDelivery(req dto.CreateDeliveryRequest) (*entity.Delivery, error) {
	delivery := &entity.Delivery{
		TrackingID:      req.TrackingID,
		NoResi:          req.NoResi,
		DeliveryAddress: req.DeliveryAddress,
		CourierName:     req.CourierName,
		CourierPhone:    req.CourierPhone,
		Status:          "OUT_FOR_DELIVERY",
	}

	if err := s.repo.Create(delivery); err != nil {
		return nil, err
	}

	return delivery, nil
}

func (s *DeliveryService) GetAllDeliveries() ([]entity.Delivery, error) {
	return s.repo.FindAll()
}

func (s *DeliveryService) GetDeliveryByID(id int) (*entity.Delivery, error) {
	return s.repo.FindByID(id)
}

func (s *DeliveryService) GetDeliveryByTrackingID(trackingID string) (*entity.Delivery, error) {
	return s.repo.FindByTrackingID(trackingID)
}

func (s *DeliveryService) GetDeliveryByNoResi(noResi string) (*entity.Delivery, error) {
	return s.repo.FindByNoResi(noResi)
}

// UpdateDeliveryStatus - update status; jika DELIVERED set delivered_at dan publish event
func (s *DeliveryService) UpdateDeliveryStatus(id int, status string) (*entity.Delivery, error) {
	delivery, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("delivery tidak ditemukan")
	}

	if status == "DELIVERED" {
		if err := s.repo.MarkDelivered(id); err != nil {
			return nil, err
		}
		delivery.Status = "DELIVERED"

		if s.publisher != nil {
			_ = s.publisher.PublishDeliveryCompleted(delivery.TrackingID, delivery.NoResi)
		}
	} else {
		if err := s.repo.UpdateStatus(id, status); err != nil {
			return nil, err
		}
		delivery.Status = status
	}

	return delivery, nil
}