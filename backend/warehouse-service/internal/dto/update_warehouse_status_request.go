package dto

type UpdateWarehouseStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
