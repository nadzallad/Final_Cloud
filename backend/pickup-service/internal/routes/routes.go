package routes

import (
	"pickup-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, pickupHandler *handler.PickupHandler) {
	api := router.Group("/api")
	{
		api.POST("/pickups", pickupHandler.CreatePickup)
		api.GET("/pickups", pickupHandler.GetPickups)
		api.GET("/pickups/:id", pickupHandler.GetPickupByID)
		api.GET("/pickups/by-tracking/:trackingNumber", pickupHandler.GetPickupByTrackingNumber)
		api.PATCH("/pickups/:id/status", pickupHandler.UpdatePickupStatus)
	}
}
