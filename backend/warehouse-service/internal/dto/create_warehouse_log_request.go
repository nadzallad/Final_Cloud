package dto

type CreateWarehouseLogRequest struct {
	UserID         int    `json:"user_id"`
	TrackingNumber string `json:"tracking_number"`
	ItemName       string `json:"item_name"`
	Stock          int    `json:"stock"`
}
