package routes

import (
	"delivery-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handler.DeliveryHandler) {
	api := router.Group("/api")
	{
		api.POST("/deliveries", h.CreateDelivery)
		api.GET("/deliveries", h.GetDeliveries)

		// Spesifik dulu sebelum wildcard :id
		api.GET("/deliveries/by-tracking/:trackingID", h.GetDeliveryByTrackingID)
		api.GET("/deliveries/by-resi/:noResi", h.GetDeliveryByNoResi)

		// Wildcard belakang
		api.GET("/deliveries/:id", h.GetDeliveryByID)
		api.PATCH("/deliveries/:id/status", h.UpdateDeliveryStatus)
	}
}
