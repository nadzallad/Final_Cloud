package repository

import (
	"warehouse-service/internal/entity"

	"gorm.io/gorm"
)

type WarehouseRepository struct {
	DB *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	return &WarehouseRepository{DB: db}
}

func (r *WarehouseRepository) Create(log *entity.WarehouseLog) error {
	return r.DB.Create(log).Error
}

func (r *WarehouseRepository) FindAll() ([]entity.WarehouseLog, error) {
	var logs []entity.WarehouseLog
	err := r.DB.Order("created_at desc").Find(&logs).Error
	return logs, err
}

func (r *WarehouseRepository) FindByID(warehouseID int) (*entity.WarehouseLog, error) {
	var log entity.WarehouseLog
	err := r.DB.Where("warehouse_id = ?", warehouseID).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *WarehouseRepository) FindByTrackingNumber(trackingNumber string) (*entity.WarehouseLog, error) {
	var log entity.WarehouseLog
	err := r.DB.Where("tracking_number = ?", trackingNumber).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *WarehouseRepository) UpdateStatus(warehouseID int, status string) error {
	return r.DB.Model(&entity.WarehouseLog{}).
		Where("warehouse_id = ?", warehouseID).
		Update("status", status).Error
}
