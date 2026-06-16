package main

import (
	"log"

	"shipment-service/internal/config"
	"shipment-service/internal/entity"
	"shipment-service/internal/handler"
	"shipment-service/internal/rabbitmq"
	"shipment-service/internal/repository"
	"shipment-service/internal/routes"
	"shipment-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db := config.ConnectDB()
	log.Println("DB connected!")

	db.AutoMigrate(&entity.Shipment{})

	shipmentRepo := repository.NewShipmentRepository(db)

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
			ch.QueueDeclare("pickup.completed", true, false, false, false, nil)
			ch.QueueDeclare("shipment.delivered", true, false, false, false, nil)
			ch.QueueDeclare("shipment.created", true, false, false, false, nil)

			publisher = &rabbitmq.Publisher{Channel: ch}

			shipmentSvc := service.NewShipmentService(shipmentRepo, publisher)

			// Consume pickup.completed → auto create shipment
			rabbitmq.ConsumePickupCompleted(ch, shipmentSvc)
			log.Println("RabbitMQ consumer aktif!")
		}
	}

	shipmentSvc := service.NewShipmentService(shipmentRepo, publisher)
	shipmentHandler := handler.NewShipmentHandler(shipmentSvc)

	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router, shipmentHandler)

	log.Println("Shipment Service running on :8085")
	router.Run(":8085")
}