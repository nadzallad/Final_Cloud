package dto

type UpdatePickupStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
