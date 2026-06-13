package handler

import (
	"net/http"

	"payment-service/internal/dto"
	"payment-service/internal/service"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	Service *service.PaymentService
}

func (h *PaymentHandler) Create(
	c *gin.Context,
) {

	var req dto.CreatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.Service.CreatePayment(req)

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
		gin.H{
			"message": "payment created",
		},
	)
}