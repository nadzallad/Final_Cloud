package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"warehouse-service/internal/dto"
	"warehouse-service/internal/entity"
	"warehouse-service/internal/rabbitmq"
	"warehouse-service/internal/repository"
)

type WarehouseService struct {
	repo      *repository.WarehouseRepository
	publisher *rabbitmq.Publisher
}

func NewWarehouseService(
	repo *repository.WarehouseRepository,
	publisher *rabbitmq.Publisher,
) *WarehouseService {
	return &WarehouseService{
		repo:      repo,
		publisher: publisher,
	}
}

func pickupServiceURL() string {
	url := os.Getenv("PICKUP_SERVICE_URL")
	if url == "" {
		url = "http://pickup-service:8083"
	}
	return url
}

func paymentServiceURL() string {
	url := os.Getenv("PAYMENT_SERVICE_URL")
	if url == "" {
		url = "http://payment-service:8082"
	}
	return url
}

func orderServiceURL() string {
	url := os.Getenv("ORDER_SERVICE_URL")
	if url == "" {
		url = "http://order-service:8081"
	}
	return url
}

type pickupResponse struct {
	UserID         int    `json:"user_id"`
	TrackingNumber string `json:"tracking_number"`
}

type orderByResiResponse struct {
	OrderID  int    `json:"order_id"`
	ItemName string `json:"item_name"`
}

// CreateLogFromPickup dipanggil saat menerima event pickup.completed.
// Membuat record warehouse_logs baru dengan status IN_WAREHOUSE dan stock default 1.
func (s *WarehouseService) CreateLogFromPickup(trackingNumber string) error {

	// Hindari duplikat jika event diterima lebih dari sekali
	existing, _ := s.repo.FindByTrackingNumber(trackingNumber)
	if existing != nil {
		return nil
	}

	userID := fetchUserIDFromPickup(trackingNumber)
	itemName := fetchItemNameByResi(trackingNumber)

	log := &entity.WarehouseLog{
		UserID:         userID,
		TrackingNumber: trackingNumber,
		ItemName:       itemName,
		Stock:          1,
		Status:         "IN_WAREHOUSE",
	}

	return s.repo.Create(log)
}

// CreateLog membuat warehouse log secara manual (REST API biasa)
func (s *WarehouseService) CreateLog(req dto.CreateWarehouseLogRequest) (*entity.WarehouseLog, error) {
	stock := req.Stock
	if stock <= 0 {
		stock = 1
	}

	log := &entity.WarehouseLog{
		UserID:         req.UserID,
		TrackingNumber: req.TrackingNumber,
		ItemName:       req.ItemName,
		Stock:          stock,
		Status:         "IN_WAREHOUSE",
	}

	if err := s.repo.Create(log); err != nil {
		return nil, err
	}

	return log, nil
}

func (s *WarehouseService) GetAllLogs() ([]entity.WarehouseLog, error) {
	return s.repo.FindAll()
}

// OverviewResponse menggabungkan data dari order, payment, pickup, dan warehouse
// supaya warehouse-service bisa berfungsi sebagai "command center" gudang.
type OverviewResponse struct {
	Orders         []map[string]interface{} `json:"orders"`
	Payments       []map[string]interface{} `json:"payments"`
	Pickups        []map[string]interface{} `json:"pickups"`
	WarehouseLogs  []entity.WarehouseLog     `json:"warehouse_logs"`
}

// GetOverview mengambil data ringkasan dari order-service, payment-service,
// dan pickup-service, digabung dengan data warehouse_logs milik service ini sendiri.
// Jika salah satu service tujuan tidak bisa dihubungi atau formatnya tidak sesuai,
// list untuk service itu dikembalikan sebagai array kosong (tidak menggagalkan request).
func (s *WarehouseService) GetOverview() (*OverviewResponse, error) {
	warehouseLogs, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	overview := &OverviewResponse{
		Orders:        fetchListAsMaps(fmt.Sprintf("%s/api/orders", orderServiceURL())),
		Payments:      fetchListAsMaps(fmt.Sprintf("%s/payments", paymentServiceURL())),
		Pickups:       fetchListAsMaps(fmt.Sprintf("%s/api/pickups", pickupServiceURL())),
		WarehouseLogs: warehouseLogs,
	}

	return overview, nil
}

// fetchListAsMaps melakukan GET ke url dan mencoba decode response sebagai
// array of object. Jika gagal terhubung atau response bukan array, kembalikan
// slice kosong (bukan nil) agar tetap aman di-marshal jadi JSON array di frontend.
func fetchListAsMaps(url string) []map[string]interface{} {
	result := []map[string]interface{}{}

	resp, err := http.Get(url)
	if err != nil {
		return result
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []map[string]interface{}{}
	}

	return result
}

func (s *WarehouseService) GetLogByID(warehouseID int) (*entity.WarehouseLog, error) {
	return s.repo.FindByID(warehouseID)
}

// UpdateLogStatus mengubah status warehouse log. Jika status baru adalah
// OUT_FOR_SHIPMENT, service ini akan publish event "warehouse.completed"
// agar shipment-service bisa membuat record shipments.
func (s *WarehouseService) UpdateLogStatus(warehouseID int, status string) (*entity.WarehouseLog, error) {
	log, err := s.repo.FindByID(warehouseID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.UpdateStatus(warehouseID, status); err != nil {
		return nil, err
	}

	log.Status = status

	if status == "OUT_FOR_SHIPMENT" && s.publisher != nil {
		if err := s.publisher.PublishWarehouseCompleted(log.TrackingNumber); err != nil {
			return nil, fmt.Errorf("gagal publish warehouse.completed: %v", err)
		}
	}

	return log, nil
}

// fetchUserIDFromPickup mengambil user_id dari pickup-service berdasarkan tracking_number.
// Jika gagal, kembalikan 0.
func fetchUserIDFromPickup(trackingNumber string) int {
	url := fmt.Sprintf("%s/api/pickups/by-tracking/%s", pickupServiceURL(), trackingNumber)

	resp, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0
	}

	var data pickupResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0
	}

	return data.UserID
}

// fetchItemNameByResi mengambil item_name dari order-service berdasarkan no_resi.
// Jika endpoint belum tersedia, kembalikan string kosong.
func fetchItemNameByResi(noResi string) string {
	url := fmt.Sprintf("%s/api/orders/by-resi/%s", orderServiceURL(), noResi)

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ""
	}

	var data orderByResiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return ""
	}

	return data.ItemName
}
