package handler

import (
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
