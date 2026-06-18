package handler

import (
	"net/http"
	"strconv"

	"notification-service/internal/dto"
	"notification-service/internal/entity"
	"notification-service/internal/service"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"fmt"
)

type NotificationHandler struct {
	service *service.NotificationService
}

func NewNotificationHandler(
	service *service.NotificationService,
) *NotificationHandler {

	return &NotificationHandler{
		service: service,
	}
}

func (h *NotificationHandler) CreateNotification(
	c *gin.Context,
) {

	var req entity.Notification

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	err := h.service.CreateNotification(req)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		dto.NotificationResponse{
			Status: "Notification created successfully",
		},
	)
}

func (h *NotificationHandler) GetNotificationsByOrderID(
	c *gin.Context,
) {
	orderID := c.Param("order_id")

	notifications, err :=
		h.service.GetNotificationsByOrderID(orderID)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		notifications,
	)
}

func (h *NotificationHandler) GetNotifications(
	c *gin.Context,
) {

	notifications, err :=
		h.service.GetAllNotifications()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		notifications,
	)
}

func (h *NotificationHandler) NotificationStream(
	c *gin.Context,
) {

	c.Writer.Header().Set(
		"Content-Type",
		"text/event-stream",
	)

	c.Writer.Header().Set(
		"Cache-Control",
		"no-cache",
	)

	c.Writer.Header().Set(
		"Connection",
		"keep-alive",
	)

	ch := make(chan entity.Notification)

	service.Clients[ch] = true

	defer func() {
		delete(service.Clients, ch)
		close(ch)
	}()

	for {
		notif := <-ch

		data, _ := json.Marshal(notif)

		fmt.Fprintf(
			c.Writer,
			"data: %s\n\n",
			data,
		)

		c.Writer.Flush()
	}
}

func (h *NotificationHandler) GetNotificationsByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	notifications, err := h.service.GetNotificationsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}


