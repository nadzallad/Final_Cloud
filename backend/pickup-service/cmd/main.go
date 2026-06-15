package main

import (
	"log"

	"pickup-service/internal/config"
	"pickup-service/internal/entity"
	"pickup-service/internal/handler"
	"pickup-service/internal/rabbitmq"
	"pickup-service/internal/repository"
	"pickup-service/internal/routes"
	"pickup-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db := config.ConnectDB()
	log.Println("DB connected!")

	db.AutoMigrate(&entity.Pickup{})

	pickupRepo := repository.NewPickupRepository(db)

	var publisher *rabbitmq.Publisher

	// RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Println("Warning: RabbitMQ tidak tersedia:", err)
	} else {
		ch, err := conn.Channel()
		if err != nil {
			log.Println("Warning: gagal buat channel RabbitMQ:", err)
		} else {
			ch.QueueDeclare("payment.success", true, false, false, false, nil)
			ch.QueueDeclare("pickup.completed", true, false, false, false, nil)

			publisher = &rabbitmq.Publisher{Channel: ch}

			pickupService := service.NewPickupService(pickupRepo, publisher)

			rabbitmq.ConsumePaymentSuccess(ch, pickupService)
			log.Println("RabbitMQ consumer aktif!")
		}
	}

	pickupService := service.NewPickupService(pickupRepo, publisher)
	pickupHandler := handler.NewPickupHandler(pickupService)

	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router, pickupHandler)

	log.Println("Pickup Service running on :8083")
	router.Run(":8083")
}
