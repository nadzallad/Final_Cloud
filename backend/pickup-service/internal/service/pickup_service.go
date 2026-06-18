package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"pickup-service/internal/dto"
	"pickup-service/internal/entity"
	"pickup-service/internal/rabbitmq"
	"pickup-service/internal/repository"
)

type PickupService struct {
	repo      *repository.PickupRepository
	publisher *rabbitmq.Publisher
}

func NewPickupService(
	repo *repository.PickupRepository,
	publisher *rabbitmq.Publisher,
) *PickupService {
	return &PickupService{
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

func trackingServiceURL() string {
	url := os.Getenv("TRACKING_SERVICE_URL")
	if url == "" {
		url = "http://tracking-service:8087"
	}
	return url
}

type trackingPayload struct {
	TrackingID      string `json:"tracking_id"`
	NoResi          string `json:"no_resi"`
	Status          string `json:"status"`
	CurrentLocation string `json:"location"`
	Note            string `json:"note"`
}

func sendTrackingUpdate(noResi string, status string, note string) {
	payload := trackingPayload{
		TrackingID:      noResi,
		NoResi:          noResi,
		Status:          status,
		CurrentLocation: "Gudang Asal",
		Note:            note,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.Println("Warning: gagal marshal payload tracking:", err)
		return
	}

	url := fmt.Sprintf("%s/tracking", trackingServiceURL())

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Warning: gagal mengirim update ke tracking-service:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		log.Println("Warning: tracking-service mengembalikan status", resp.StatusCode)
	}
}

type resiResponse struct {
	OrderID int    `json:"order_id"`
	NoResi  string `json:"no_resi"`
}

type orderDetailResponse struct {
	OrderID  int     `json:"order_id"`
	UserID   int     `json:"user_id"`
	WeightKg float64 `json:"weight_kg"`
}

func (s *PickupService) CreatePickupFromPayment(orderID string) error {
	noResi, err := fetchNoResi(orderID)
	if err != nil {
		return fmt.Errorf("gagal mengambil no_resi dari order-service: %v", err)
	}

	existing, _ := s.repo.FindByTrackingNumber(noResi)
	if existing != nil {
		return nil
	}

	orderIDInt, _ := strconv.Atoi(orderID)
	userID, weightKg := fetchOrderDetail(orderID)

	pickup := &entity.Pickup{
		OrderID:        orderIDInt,
		UserID:         userID,
		TrackingNumber: noResi,
		PaymentStatus:  "PAID",
		WeightKg:       weightKg,
		Status:         "WAITING_PICKUP",
	}

	if err := s.repo.Create(pickup); err != nil {
		return err
	}

	go sendTrackingUpdate(noResi, "WAITING_PICKUP", "Paket menunggu diambil kurir")

	return nil
}

func (s *PickupService) CreatePickup(req dto.CreatePickupRequest) (*entity.Pickup, error) {
	pickup := &entity.Pickup{
		UserID:         req.UserID,
		TrackingNumber: req.TrackingNumber,
		PaymentStatus:  req.PaymentStatus,
		WeightKg:       req.WeightKg,
		Status:         "WAITING_PICKUP",
	}

	if err := s.repo.Create(pickup); err != nil {
		return nil, err
	}

	go sendTrackingUpdate(req.TrackingNumber, "WAITING_PICKUP", "Paket menunggu diambil kurir")

	return pickup, nil
}

func (s *PickupService) GetAllPickups() ([]entity.Pickup, error) {
	return s.repo.FindAll()
}

func (s *PickupService) GetPickupByID(pickupID int) (*entity.Pickup, error) {
	return s.repo.FindByID(pickupID)
}

func (s *PickupService) GetPickupByTrackingNumber(trackingNumber string) (*entity.Pickup, error) {
	return s.repo.FindByTrackingNumber(trackingNumber)
}

func (s *PickupService) UpdatePickupStatus(pickupID int, status string) (*entity.Pickup, error) {
	pickup, err := s.repo.FindByID(pickupID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.UpdateStatus(pickupID, status); err != nil {
		return nil, err
	}

	pickup.Status = status

	if status == "PICKED_UP" && s.publisher != nil {
		orderIDStr := strconv.Itoa(pickup.OrderID)
		if err := s.publisher.PublishPickupCompleted(pickup.TrackingNumber, orderIDStr); err != nil {
			return nil, fmt.Errorf("gagal publish pickup.completed: %v", err)
		}
	}

	if status == "PICKED_UP" {
		go sendTrackingUpdate(pickup.TrackingNumber, "PICKED_UP", "Paket telah diambil kurir dari pengirim")
	}

	return pickup, nil
}

func fetchNoResi(orderID string) (string, error) {
	url := fmt.Sprintf("%s/api/orders/%s/resi", orderServiceURL(), orderID)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("order-service returned status %d", resp.StatusCode)
	}

	var data resiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.NoResi == "" {
		return "", fmt.Errorf("no_resi kosong")
	}

	return data.NoResi, nil
}

func fetchOrderDetail(orderID string) (userID int, weightKg float64) {
	url := fmt.Sprintf("%s/api/orders/%s", orderServiceURL(), orderID)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0
	}

	var data orderDetailResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, 0
	}

	return data.UserID, data.WeightKg
}
