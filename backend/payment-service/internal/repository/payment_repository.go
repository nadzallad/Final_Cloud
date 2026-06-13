package repository

import (
	"context"

	"payment-service/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	DB *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

func (r *PaymentRepository) Create(
	payment entity.Payment,
) error {

	query := `
	INSERT INTO payments
	(
		payment_id,
		order_id,
		payment_method,
		total,
		discount,
		admin_fee,
		status
	)
	VALUES
	($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		payment.PaymentID,
		payment.OrderID,
		payment.PaymentMethod,
		payment.Total,
		payment.Discount,
		payment.AdminFee,
		payment.Status,
	)

	return err
}