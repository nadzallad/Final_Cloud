package dto

type CreateWarehouseLogRequest struct {
	UserID         int    `json:"user_id"`
	TrackingNumber string `json:"tracking_number" binding:"required"`
	ItemName       string `json:"item_name"`
	Stock          int    `json:"stock"`
}

type UpdateWarehouseStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type WarehouseLogResponse struct {
	WarehouseID    int    `json:"warehouse_id"`
	UserID         int    `json:"user_id"`
	TrackingNumber string `json:"tracking_number"`
	ItemName       string `json:"item_name"`
	Stock          int    `json:"stock"`
	Status         string `json:"status"`
}
