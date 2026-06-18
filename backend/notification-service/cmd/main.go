package main

import (
	"log"
	"time"

	"notification-service/internal/config"
	"notification-service/internal/handler"
	"notification-service/internal/rabbitmq"
	"notification-service/internal/repository"
	"notification-service/internal/routes"
	"notification-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
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

	// RabbitMQ - connect dengan retry
	var conn *amqp.Connection
	var err error
	for i := 0; i < 5; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			break
		}
		log.Printf("RabbitMQ belum siap, retry %d/5...", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Println("Warning: RabbitMQ tidak tersedia:", err)
	} else {
		ch, err := conn.Channel()
		if err != nil {
			log.Println("Warning: gagal buat channel RabbitMQ:", err)
		} else {
			rabbitmq.ConsumeDeliveryCompleted(ch, notificationService)
			log.Println("RabbitMQ consumer delivery.completed aktif!")
		}
	}

	router.Run(":8088")

}
