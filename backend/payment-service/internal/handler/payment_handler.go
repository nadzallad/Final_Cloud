package handler

import (
	"log"
	"net/http"
	"payment-service/internal/dto"
	"payment-service/internal/service"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(
	service *service.PaymentService,
) *PaymentHandler {

	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) CreatePayment(
	c *gin.Context,
) {

	var req dto.CreatePaymentRequest

	if err :=
		c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	payment, err :=
		h.service.CreatePayment(req)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		payment,
	)
}

func (h *PaymentHandler) Pay(
	c *gin.Context,
) {

	id := c.Param("id")

	payment, err :=
		h.service.MarkAsPaid(id)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		payment,
	)
}

func (h *PaymentHandler) Notification(
	c *gin.Context,
) {

	var notification map[string]interface{}

	log.Println("=== MIDTRANS CALLBACK MASUK ===")

	if err := c.ShouldBindJSON(&notification); err != nil {

		log.Printf("ShouldBindJSON error: %v", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	log.Printf(
		"Payload: %+v\n",
		notification,
	)

	orderID, ok :=
		notification["order_id"].(string)

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "order_id tidak ditemukan",
			},
		)

		return
	}

	transactionStatus, ok :=
		notification["transaction_status"].(string)

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "transaction_status tidak ditemukan",
			},
		)

		return
	}

	log.Printf("OrderID: %s", orderID)
	log.Printf("Status : %s", transactionStatus)

	if transactionStatus == "settlement" {

		_, err :=
			h.service.MarkAsPaid(
				orderID,
			)

		if err != nil {

			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": err.Error(),
				},
			)

			return
		}
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "notification received",
		},
	)
}

func (h *PaymentHandler) CheckStatus(
	c *gin.Context,
) {

	orderID := c.Param("orderId")

	status, err := service.CheckTransaction(orderID)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	
	if status == "settlement" {
		h.service.MarkAsPaid(orderID)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": status,
		},
	)
}
