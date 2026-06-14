package main

import (
	"log"
	"order-service/internal/config"
	"order-service/internal/entity"
	"order-service/internal/handler"
	"order-service/internal/rabbitmq"
	"order-service/internal/repository"
	"order-service/internal/routes"
	"order-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db := config.ConnectDB()
	log.Println("DB connected!")

	db.AutoMigrate(&entity.Order{})

	orderRepo := repository.NewOrderRepository(db)
	cityRepo := repository.NewCityRepository(db)
	orderService := service.NewOrderService(orderRepo, cityRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	// RabbitMQ consumer
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Warning: RabbitMQ tidak tersedia:", err)
	} else {
		ch, err := conn.Channel()
		if err != nil {
			log.Println("Warning: gagal buat channel RabbitMQ:", err)
		} else {
			ch.QueueDeclare("payment.success", true, false, false, false, nil)
			rabbitmq.ConsumePaymentSuccess(ch, orderService)
			log.Println("RabbitMQ consumer aktif!")
		}
	}

	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router, orderHandler)

	log.Println("Order Service running on :8081")
	router.Run(":8081")
}