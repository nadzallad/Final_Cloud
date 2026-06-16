package repository

import (
	"shipment-service/internal/entity"

	"gorm.io/gorm"
)

type ShipmentRepository struct {
	DB *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) *ShipmentRepository {
	return &ShipmentRepository{DB: db}
}

func (r *ShipmentRepository) Create(shipment *entity.Shipment) error {
	return r.DB.Create(shipment).Error
}

func (r *ShipmentRepository) FindAll() ([]entity.Shipment, error) {
	var shipments []entity.Shipment
	err := r.DB.Order("created_at desc").Find(&shipments).Error
	return shipments, err
}

func (r *ShipmentRepository) FindByID(id int) (*entity.Shipment, error) {
	var shipment entity.Shipment
	err := r.DB.Where("shipment_id = ?", id).First(&shipment).Error
	if err != nil {
		return nil, err
	}
	return &shipment, nil
}

func (r *ShipmentRepository) FindByNoResi(noResi string) (*entity.Shipment, error) {
	var shipment entity.Shipment
	err := r.DB.Where("no_resi = ?", noResi).First(&shipment).Error
	if err != nil {
		return nil, err
	}
	return &shipment, nil
}

func (r *ShipmentRepository) FindByTrackingID(trackingID string) (*entity.Shipment, error) {
	var shipment entity.Shipment
	err := r.DB.Where("tracking_id = ?", trackingID).First(&shipment).Error
	if err != nil {
		return nil, err
	}
	return &shipment, nil
}

func (r *ShipmentRepository) UpdateStatus(id int, status string, currentLocation string) error {
	updates := map[string]interface{}{"status": status}
	if currentLocation != "" {
		updates["current_location"] = currentLocation
	}
	return r.DB.Model(&entity.Shipment{}).
		Where("shipment_id = ?", id).
		Updates(updates).Error
}