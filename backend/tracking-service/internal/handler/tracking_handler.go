package handler

import (
	"net/http"

	"tracking-service/internal/dto"
	"tracking-service/internal/entity"
	"tracking-service/internal/service"

	"github.com/gin-gonic/gin"
)

type TrackingHandler struct {
	service *service.TrackingService
}

func NewTrackingHandler(
	service *service.TrackingService,
) *TrackingHandler {

	return &TrackingHandler{
		service: service,
	}
}

func (h *TrackingHandler) CreateTracking(
	c *gin.Context,
) {

	var req entity.Tracking

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	err := h.service.CreateTracking(req)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		dto.TrackingResponse{
			Status: req.Status,
		},
	)
}

func (h *TrackingHandler) GetTrackingByResi(
	c *gin.Context,
) {

	noResi := c.Param("no_resi")

	tracking, err := h.service.GetTrackingByResi(noResi)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		tracking,
	)
}