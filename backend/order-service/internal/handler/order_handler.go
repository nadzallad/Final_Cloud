package handler

import (
	"net/http"

	"order-service/internal/dto"
	"order-service/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(
	service *service.OrderService,
) *OrderHandler {

	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(
	c *gin.Context,
) {

	var req dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	order, err :=
		h.service.CreateOrder(req)

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
		order,
	)
}

func (h *OrderHandler) GetOrders(
	c *gin.Context,
) {

	orders, err :=
		h.service.GetOrders()

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
		http.StatusOK,
		orders,
	)
}