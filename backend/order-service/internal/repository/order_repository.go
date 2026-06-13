package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	DB *pgxpool.Pool
}

func NewOrderRepository(
	db *pgxpool.Pool,
) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) UpdateStatus(
	orderID int,
	status string,
) error {

	query := `
	UPDATE orders
	SET status=$1
	WHERE order_id=$2
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		status,
		orderID,
	)

	return err
}