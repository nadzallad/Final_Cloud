package routes

import (
	"payment-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	paymentHandler *handler.PaymentHandler,
) {

	router.GET("/payments", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Payment Service Running",
		})
	})

	router.POST(
		"/payments",
		paymentHandler.CreatePayment,
	)
	
	router.PATCH(
		"/payments/:id/pay",
		paymentHandler.Pay,
	)
	
	router.POST(
		"/payments/notification",
		paymentHandler.Notification,
	)
	

}
