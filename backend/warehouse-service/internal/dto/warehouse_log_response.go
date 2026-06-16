package dto

type WarehouseLogResponse struct {
	WarehouseID    int    `json:"warehouse_id"`
	UserID         int    `json:"user_id"`
	TrackingNumber string `json:"tracking_number"`
	ItemName       string `json:"item_name"`
	Stock          int    `json:"stock"`
	Status         string `json:"status"`
}
