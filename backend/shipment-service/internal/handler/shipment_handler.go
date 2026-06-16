package handler

import (
	"net/http"
	"strconv"

	"shipment-service/internal/dto"
	"shipment-service/internal/service"

	"github.com/gin-gonic/gin"
)

type ShipmentHandler struct {
	service *service.ShipmentService
}

func NewShipmentHandler(service *service.ShipmentService) *ShipmentHandler {
	return &ShipmentHandler{service: service}
}

func (h *ShipmentHandler) CreateShipment(c *gin.Context) {
	var req dto.CreateShipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shipment, err := h.service.CreateShipment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, shipment)
}

func (h *ShipmentHandler) GetShipments(c *gin.Context) {
	shipments, err := h.service.GetAllShipments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shipments)
}

func (h *ShipmentHandler) GetShipmentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shipment id"})
		return
	}

	shipment, err := h.service.GetShipmentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shipment not found"})
		return
	}

	c.JSON(http.StatusOK, shipment)
}

func (h *ShipmentHandler) GetShipmentByNoResi(c *gin.Context) {
	noResi := c.Param("noResi")
	shipment, err := h.service.GetShipmentByNoResi(noResi)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shipment not found"})
		return
	}

	c.JSON(http.StatusOK, shipment)
}

func (h *ShipmentHandler) GetShipmentByTrackingID(c *gin.Context) {
	trackingID := c.Param("trackingID")
	shipment, err := h.service.GetShipmentByTrackingID(trackingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shipment not found"})
		return
	}

	c.JSON(http.StatusOK, shipment)
}

func (h *ShipmentHandler) UpdateShipmentStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shipment id"})
		return
	}

	var req dto.UpdateShipmentStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shipment, err := h.service.UpdateShipmentStatus(id, req.Status, req.CurrentLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shipment)
}