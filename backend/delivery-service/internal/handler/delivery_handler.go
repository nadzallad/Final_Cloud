package handler

import (
	"net/http"
	"strconv"

	"delivery-service/internal/dto"
	"delivery-service/internal/service"

	"github.com/gin-gonic/gin"
)

type DeliveryHandler struct {
	service *service.DeliveryService
}

func NewDeliveryHandler(service *service.DeliveryService) *DeliveryHandler {
	return &DeliveryHandler{service: service}
}

func (h *DeliveryHandler) CreateDelivery(c *gin.Context) {
	var req dto.CreateDeliveryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delivery, err := h.service.CreateDelivery(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, delivery)
}

func (h *DeliveryHandler) GetDeliveries(c *gin.Context) {
	deliveries, err := h.service.GetAllDeliveries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deliveries)
}

func (h *DeliveryHandler) GetDeliveryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid delivery id"})
		return
	}

	delivery, err := h.service.GetDeliveryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "delivery not found"})
		return
	}

	c.JSON(http.StatusOK, delivery)
}

func (h *DeliveryHandler) GetDeliveryByTrackingID(c *gin.Context) {
	trackingID := c.Param("trackingID")
	delivery, err := h.service.GetDeliveryByTrackingID(trackingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "delivery not found"})
		return
	}

	c.JSON(http.StatusOK, delivery)
}

func (h *DeliveryHandler) GetDeliveryByNoResi(c *gin.Context) {
	noResi := c.Param("noResi")
	delivery, err := h.service.GetDeliveryByNoResi(noResi)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "delivery not found"})
		return
	}

	c.JSON(http.StatusOK, delivery)
}

func (h *DeliveryHandler) UpdateDeliveryStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid delivery id"})
		return
	}

	var req dto.UpdateDeliveryStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delivery, err := h.service.UpdateDeliveryStatus(id, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, delivery)
}