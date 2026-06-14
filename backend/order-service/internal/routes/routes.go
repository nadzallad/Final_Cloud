package routes

import (
	"order-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	orderHandler *handler.OrderHandler,
) {
	router.POST(
		"/orders",
		orderHandler.CreateOrder,
	)

	router.GET(
		"/orders",
		orderHandler.GetOrders,
	)
}