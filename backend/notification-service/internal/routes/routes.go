package routes

import (
	"notification-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	notificationHandler *handler.NotificationHandler,
) {

	router.GET(
		"/notification/stream",
		notificationHandler.NotificationStream,
	)

	router.GET(
		"/notification/by-user/:user_id",
		notificationHandler.GetNotificationsByUserID,
	)
	
	router.POST(
		"/notification",
		notificationHandler.CreateNotification,
	)

	router.GET(
		"/notification/:order_id",
		notificationHandler.GetNotificationsByOrderID,
	)

	router.GET(
		"/notification",
		notificationHandler.GetNotifications,
	)
}	
