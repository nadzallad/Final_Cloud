package routes

import (
	"shipment-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handler.ShipmentHandler) {
	api := router.Group("/api")
	{
		api.POST("/shipments", h.CreateShipment)
		api.GET("/shipments", h.GetShipments)
		api.GET("/shipments/by-resi/:noResi", h.GetShipmentByNoResi)
		api.GET("/shipments/by-tracking/:trackingID", h.GetShipmentByTrackingID)
		api.GET("/shipments/:id", h.GetShipmentByID)
		api.PATCH("/shipments/by-resi/:noResi/status", h.UpdateShipmentStatus)
	}
}
