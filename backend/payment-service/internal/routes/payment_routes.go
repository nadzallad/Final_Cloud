package routes

import (
	"payment-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	handler *handler.PaymentHandler,
) {

	payment := router.Group("/payments")

	{
		payment.POST("", handler.Create)
	}
}