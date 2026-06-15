package repository

import (
	"payment-service/internal/entity"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(
	db *gorm.DB,
) *PaymentRepository {

	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) Create(
	payment *entity.Payment,
) error {

	return r.db.Create(payment).Error
}

func (r *PaymentRepository) FindByID(
	id string,
) (*entity.Payment, error) {

	var payment entity.Payment

	err := r.db.
		Where("payment_id = ?", id).
		First(&payment).Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) Update(
	payment *entity.Payment,
) error {

	return r.db.Save(payment).Error
}

func (r *PaymentRepository) FindByOrderID(
	orderID int,
) (*entity.Payment, error) {

	var payment entity.Payment

	err := r.db.
		Where(
			"order_id = ?",
			orderID,
		).
		First(&payment).Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}