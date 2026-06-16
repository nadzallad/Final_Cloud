package main

import (
	"time"

	"notification-service/internal/config"
	"notification-service/internal/handler"
	"notification-service/internal/repository"
	"notification-service/internal/routes"
	"notification-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	notificationRepo :=
		repository.NewNotificationRepository(
			config.NotificationCollection,
		)

	notificationService :=
		service.NewNotificationService(
			notificationRepo,
		)

	notificationHandler :=
		handler.NewNotificationHandler(
			notificationService,
		)

	routes.SetupRoutes(
		router,
		notificationHandler,
	)

	router.Run(":8088")
}