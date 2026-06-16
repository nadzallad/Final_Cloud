package routes

import (
	"fmt"
    "net/http"
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
	router.GET("/", func(c *gin.Context) {

    orderID := c.Query("order_id")
    status := c.Query("transaction_status")

    c.Redirect(
			http.StatusFound,
			fmt.Sprintf(
				"http://localhost:5173/dashboard?order_id=%s&status=%s",
				orderID,
				status,
			),
		)
	})
		
 
}
