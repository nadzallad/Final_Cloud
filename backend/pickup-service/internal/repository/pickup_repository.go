package repository

import (
	"pickup-service/internal/entity"

	"gorm.io/gorm"
)

type PickupRepository struct {
	DB *gorm.DB
}

func NewPickupRepository(db *gorm.DB) *PickupRepository {
	return &PickupRepository{DB: db}
}

func (r *PickupRepository) Create(pickup *entity.Pickup) error {
	return r.DB.Create(pickup).Error
}

func (r *PickupRepository) FindAll() ([]entity.Pickup, error) {
	var pickups []entity.Pickup
	err := r.DB.Order("created_at desc").Find(&pickups).Error
	return pickups, err
}

func (r *PickupRepository) FindByID(pickupID int) (*entity.Pickup, error) {
	var pickup entity.Pickup
	err := r.DB.Where("pickup_id = ?", pickupID).First(&pickup).Error
	if err != nil {
		return nil, err
	}
	return &pickup, nil
}

func (r *PickupRepository) FindByTrackingNumber(trackingNumber string) (*entity.Pickup, error) {
	var pickup entity.Pickup
	err := r.DB.Where("tracking_number = ?", trackingNumber).First(&pickup).Error
	if err != nil {
		return nil, err
	}
	return &pickup, nil
}

func (r *PickupRepository) UpdateStatus(pickupID int, status string) error {
	return r.DB.Model(&entity.Pickup{}).
		Where("pickup_id = ?", pickupID).
		Update("status", status).Error
}
