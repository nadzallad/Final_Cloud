package handler

import (
	"net/http"
	"strconv"

	"pickup-service/internal/dto"
	"pickup-service/internal/service"

	"github.com/gin-gonic/gin"
)

type PickupHandler struct {
	service *service.PickupService
}

func NewPickupHandler(service *service.PickupService) *PickupHandler {
	return &PickupHandler{
		service: service,
	}
}

func (h *PickupHandler) CreatePickup(c *gin.Context) {
	var req dto.CreatePickupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pickup, err := h.service.CreatePickup(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pickup)
}

func (h *PickupHandler) GetPickups(c *gin.Context) {
	pickups, err := h.service.GetAllPickups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pickups)
}

func (h *PickupHandler) GetPickupByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pickup id"})
		return
	}

	pickup, err := h.service.GetPickupByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pickup not found"})
		return
	}

	c.JSON(http.StatusOK, pickup)
}

func (h *PickupHandler) GetPickupByTrackingNumber(c *gin.Context) {
	trackingNumber := c.Param("trackingNumber")

	pickup, err := h.service.GetPickupByTrackingNumber(trackingNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pickup not found"})
		return
	}

	c.JSON(http.StatusOK, pickup)
}

func (h *PickupHandler) UpdatePickupStatus(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pickup id"})
		return
	}

	var req dto.UpdatePickupStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pickup, err := h.service.UpdatePickupStatus(id, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pickup)
}
