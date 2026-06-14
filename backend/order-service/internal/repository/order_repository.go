package repository

import (
	"order-service/internal/entity"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Create(order *entity.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.DB.Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) FindByID(orderID int) (*entity.Order, error) {
	var order entity.Order
	err := r.DB.Where("order_id = ?", orderID).First(&order).Error
	return &order, err
}

func (r *OrderRepository) UpdateStatus(orderID int, status string) error {
	return r.DB.Model(&entity.Order{}).
		Where("order_id = ?", orderID).
		Update("status", status).Error
}

func (r *OrderRepository) CreateResi(orderID int, noResi string) error {
	return r.DB.Exec(
		`INSERT INTO resi (order_id, no_resi) VALUES (?, ?)`,
		orderID, noResi,
	).Error
}

func (r *OrderRepository) FindResiByOrderID(orderID int) (string, error) {
	var noResi string
	err := r.DB.Raw(
		`SELECT no_resi FROM resi WHERE order_id = ? AND status = 'ACTIVE' LIMIT 1`,
		orderID,
	).Scan(&noResi).Error
	return noResi, err
}