package routes

import (
	"order-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, orderHandler *handler.OrderHandler) {
	api := router.Group("/api")
	{
		api.POST("/orders", orderHandler.CreateOrder)
		api.GET("/orders", orderHandler.GetOrders)
		api.POST("/orders/:id/confirm-payment", orderHandler.ConfirmPayment)
	}
}