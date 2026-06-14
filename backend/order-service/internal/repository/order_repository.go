package repository

import (
	"context"
	"order-service/internal/entity"
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
	orderID string,
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


func (r *OrderRepository) FindAll() (
	[]entity.Order,
	error,
) {

	query := `
	SELECT
		order_id,
		sender_name,
		receiver_name,
		weight_kg,
		shipping_cost,
		total_price,
		status,
		created_at
	FROM orders
	ORDER BY order_id DESC
	`

	rows, err := r.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []entity.Order

	for rows.Next() {

		var order entity.Order

		err := rows.Scan(
			&order.OrderID,
			&order.SenderName,
			&order.ReceiverName,
			&order.WeightKg,
			&order.ShippingCost,
			&order.TotalPrice,
			&order.Status,
			&order.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		orders = append(
			orders,
			order,
		)
	}

	return orders, nil
}