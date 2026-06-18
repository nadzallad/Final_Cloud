package repository

import (
	"delivery-service/internal/entity"

	"gorm.io/gorm"
)

type DeliveryRepository struct {
	DB *gorm.DB
}

func NewDeliveryRepository(db *gorm.DB) *DeliveryRepository {
	return &DeliveryRepository{DB: db}
}

func (r *DeliveryRepository) Create(delivery *entity.Delivery) error {
	return r.DB.Create(delivery).Error
}

func (r *DeliveryRepository) FindAll() ([]entity.Delivery, error) {
	var deliveries []entity.Delivery
	err := r.DB.Order("created_at desc").Find(&deliveries).Error
	return deliveries, err
}

func (r *DeliveryRepository) FindByID(id int) (*entity.Delivery, error) {
	var delivery entity.Delivery
	err := r.DB.Where("delivery_id = ?", id).First(&delivery).Error
	if err != nil {
		return nil, err
	}
	return &delivery, nil
}

func (r *DeliveryRepository) FindByTrackingID(trackingID string) (*entity.Delivery, error) {
	var delivery entity.Delivery
	err := r.DB.Where("tracking_id = ?", trackingID).First(&delivery).Error
	if err != nil {
		return nil, err
	}
	return &delivery, nil
}

func (r *DeliveryRepository) FindByNoResi(noResi string) (*entity.Delivery, error) {
	var delivery entity.Delivery
	err := r.DB.Where("no_resi = ?", noResi).First(&delivery).Error
	if err != nil {
		return nil, err
	}
	return &delivery, nil
}

func (r *DeliveryRepository) UpdateStatus(id int, status string) error {
	return r.DB.Model(&entity.Delivery{}).
		Where("delivery_id = ?", id).
		Update("status", status).Error
}

func (r *DeliveryRepository) MarkDelivered(id int) error {
	return r.DB.Model(&entity.Delivery{}).
		Where("delivery_id = ?", id).
		Updates(map[string]interface{}{
			"status":       "DELIVERED",
			"delivered_at": gorm.Expr("CURRENT_TIMESTAMP"),
		}).Error
}