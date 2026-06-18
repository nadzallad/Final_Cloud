package main

import (
	"log"

	"delivery-service/internal/config"
	"delivery-service/internal/entity"
	"delivery-service/internal/handler"
	"delivery-service/internal/rabbitmq"
	"delivery-service/internal/repository"
	"delivery-service/internal/routes"
	"delivery-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db := config.ConnectDB()
	log.Println("DB connected!")

	db.AutoMigrate(&entity.Delivery{})

	deliveryRepo := repository.NewDeliveryRepository(db)

	var publisher *rabbitmq.Publisher

	// RabbitMQ setup
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Println("Warning: RabbitMQ tidak tersedia:", err)
	} else {
		ch, err := conn.Channel()
		if err != nil {
			log.Println("Warning: gagal buat channel RabbitMQ:", err)
		} else {
			// Declare queues
			ch.QueueDeclare("shipment.delivered", true, false, false, false, nil)
			ch.QueueDeclare("delivery.completed", true, false, false, false, nil)

			publisher = &rabbitmq.Publisher{Channel: ch}

			deliverySvc := service.NewDeliveryService(deliveryRepo, publisher)

			// Consume shipment.delivered → auto create delivery
			rabbitmq.ConsumeShipmentDelivered(ch, deliverySvc)
			log.Println("RabbitMQ consumer aktif!")
		}
	}

	deliverySvc := service.NewDeliveryService(deliveryRepo, publisher)
	deliveryHandler := handler.NewDeliveryHandler(deliverySvc)

	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router, deliveryHandler)

	log.Println("Delivery Service running on :8086")
	router.Run(":8086")
}