package main

import (
	"log"

	"payment-service/internal/config"
	"payment-service/internal/handler"
	"payment-service/internal/rabbitmq"
	"payment-service/internal/repository"
	"payment-service/internal/routes"
	"payment-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

func main() {

	config.ConnectDB()

	// RabbitMQ Connection
	conn, err := amqp091.Dial(
		"amqp://guest:guest@localhost:5672/",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	// Declare Queue
	_, err = ch.QueueDeclare(
		"payment.success",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	publisher := &rabbitmq.Publisher{
		Channel: ch,
	}

	router := gin.Default()

	router.Use(cors.Default())

	paymentRepo :=
		repository.NewPaymentRepository(
			config.DB,
		)

	paymentService :=
		service.NewPaymentService(
			paymentRepo,
			publisher,
		)

	paymentHandler :=
		handler.NewPaymentHandler(
			paymentService,
		)

	routes.SetupRoutes(
		router,
		paymentHandler,
	)

	router.Run(":8082")
}