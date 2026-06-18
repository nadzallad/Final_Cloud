package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"bytes"

	"shipment-service/internal/dto"
	"shipment-service/internal/entity"
	"shipment-service/internal/rabbitmq"
	"shipment-service/internal/repository"
)

type ShipmentService struct {
	repo      *repository.ShipmentRepository
	publisher *rabbitmq.Publisher
}

func NewShipmentService(repo *repository.ShipmentRepository, publisher *rabbitmq.Publisher) *ShipmentService {
	return &ShipmentService{
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

// orderDetailForShipment dipakai untuk ambil origin/destination city dari order-service
type orderDetailForShipment struct {
	OrderID         int    `json:"order_id"`
	OriginCity      string `json:"origin_city"`
	DestinationCity string `json:"destination_city"`
}

func fetchOrderForShipment(orderID string) (*orderDetailForShipment, error) {
	url := fmt.Sprintf("%s/api/orders/%s", orderServiceURL(), orderID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("order-service returned status %d", resp.StatusCode)
	}

	var data orderDetailForShipment
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

// CreateShipmentFromPickup dipanggil saat menerima event pickup.completed
func (s *ShipmentService) CreateShipmentFromPickup(trackingNumber string, orderID string) error {
	// Cek duplikat
	existing, _ := s.repo.FindByNoResi(trackingNumber)
	if existing != nil {
		return nil
	}

	// Default origin/destination - fallback jika order-service tidak tersedia
	originCity := "Unknown"
	destinationCity := "Unknown"

	if orderID != "" {
		if detail, err := fetchOrderForShipment(orderID); err == nil {
			if detail.OriginCity != "" {
				originCity = detail.OriginCity
			}
			if detail.DestinationCity != "" {
				destinationCity = detail.DestinationCity
			}
		}
	}

	eta := time.Now().Add(3 * 24 * time.Hour) // estimasi 3 hari

	shipment := &entity.Shipment{
		TrackingID:      orderID,
		NoResi:          trackingNumber,
		OriginCity:      originCity,
		DestinationCity: destinationCity,
		CurrentLocation: originCity,
		Status:          "IN_TRANSIT",
		ETA:             &eta,
	}

	if err := s.repo.Create(shipment); err != nil {
		return err
	}

	// Publish event shipment.created untuk tracking/notification service
	if s.publisher != nil {
		_ = s.publisher.PublishShipmentCreated(trackingNumber, orderID, originCity, destinationCity)
	}

	return nil
}

// CreateShipment - manual via REST
func (s *ShipmentService) CreateShipment(req dto.CreateShipmentRequest) (*entity.Shipment, error) {
	// Cek duplikat no_resi
	existing, _ := s.repo.FindByNoResi(req.NoResi)
	if existing != nil {
		return nil, fmt.Errorf("shipment dengan no_resi %s sudah ada", req.NoResi)
	}

	currentLoc := req.CurrentLocation
	if currentLoc == "" {
		currentLoc = req.OriginCity
	}

	shipment := &entity.Shipment{
		TrackingID:      req.TrackingID,
		NoResi:          req.NoResi,
		OriginCity:      req.OriginCity,
		DestinationCity: req.DestinationCity,
		CurrentLocation: currentLoc,
		Status:          "IN_TRANSIT",
		ETA:             req.ETA,
	}

	if err := s.repo.Create(shipment); err != nil {
		return nil, err
	}

	if s.publisher != nil {
		_ = s.publisher.PublishShipmentCreated(req.NoResi, req.TrackingID, req.OriginCity, req.DestinationCity)
	}

	return shipment, nil
}

func (s *ShipmentService) GetAllShipments() ([]entity.Shipment, error) {
	return s.repo.FindAll()
}

func (s *ShipmentService) GetShipmentByID(id int) (*entity.Shipment, error) {
	return s.repo.FindByID(id)
}

func (s *ShipmentService) GetShipmentByNoResi(noResi string) (*entity.Shipment, error) {
	return s.repo.FindByNoResi(noResi)
}

func (s *ShipmentService) GetShipmentByTrackingID(trackingID string) (*entity.Shipment, error) {
	return s.repo.FindByTrackingID(trackingID)
}

// UpdateShipmentStatus - update status dan lokasi terkini
// Jika status jadi DELIVERED, publish event shipment.delivered
func (s *ShipmentService) UpdateShipmentStatus(
	noResi string,
	status string,
	currentLocation string,
) (*entity.Shipment, error) {

	shipment, err := s.repo.FindByNoResi(noResi)
	if err != nil {
		return nil, fmt.Errorf("shipment tidak ditemukan")
	}

	if err := s.repo.UpdateStatus(shipment.ShipmentID, status, currentLocation); err != nil {
		return nil, err
	}

	shipment.Status = status
	if currentLocation != "" {
		shipment.CurrentLocation = currentLocation
	}

	tracking := map[string]interface{}{
		"no_resi": shipment.NoResi,
		"status": status,
		"location": currentLocation,
		"note": "Status diperbarui",
	}

	body, _ := json.Marshal(tracking)

	resp, err := http.Post(
		"http://localhost:8087/tracking",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		fmt.Println("Gagal kirim tracking:", err)
	} else {
		fmt.Println("Status tracking:", resp.Status)
	}

	if status == "DELIVERED" && s.publisher != nil {
		if err := s.publisher.PublishShipmentDelivered(
			shipment.NoResi,
			shipment.TrackingID,
		); err != nil {
			return nil, fmt.Errorf(
				"gagal publish shipment.delivered: %v",
				err,
			)
		}
	}

	return shipment, nil
}