package main

import (
	"log"

	"warehouse-service/internal/config"
	"warehouse-service/internal/entity"
	"warehouse-service/internal/handler"
	"warehouse-service/internal/rabbitmq"
	"warehouse-service/internal/repository"
	"warehouse-service/internal/routes"
	"warehouse-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db := config.ConnectDB()
	log.Println("DB connected!")

	db.AutoMigrate(&entity.WarehouseLog{})

	warehouseRepo := repository.NewWarehouseRepository(db)

	var publisher *rabbitmq.Publisher

	// RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@host.docker.internal:5672/")
	if err != nil {
		log.Println("Warning: RabbitMQ tidak tersedia:", err)
	} else {
		ch, err := conn.Channel()
		if err != nil {
			log.Println("Warning: gagal buat channel RabbitMQ:", err)
		} else {
			ch.QueueDeclare("pickup.completed", true, false, false, false, nil)
			ch.QueueDeclare("warehouse.completed", true, false, false, false, nil)

			publisher = &rabbitmq.Publisher{Channel: ch}

			warehouseService := service.NewWarehouseService(warehouseRepo, publisher)

			rabbitmq.ConsumePickupCompleted(ch, warehouseService)
			log.Println("RabbitMQ consumer aktif!")
		}
	}

	warehouseService := service.NewWarehouseService(warehouseRepo, publisher)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)

	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router, warehouseHandler)

	log.Println("Warehouse Service running on :8084")
	router.Run(":8084")
}
