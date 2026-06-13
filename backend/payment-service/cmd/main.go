package main

import (
	"log"

	"payment-service/internal/config"
	"payment-service/internal/handler"
	"payment-service/internal/rabbitmq"
	"payment-service/internal/repository"
	"payment-service/internal/routes"
	"payment-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

func main() {

	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	conn, err := amqp091.Dial(
		"amqp://guest:guest@rabbitmq:5672/",
	)

	if err != nil {
		log.Fatal(err)
	}

	ch, _ := conn.Channel()

	publisher := &rabbitmq.Publisher{
		Channel: ch,
	}

	repo := repository.NewPaymentRepository(db)

	svc := &service.PaymentService{
		Repo: repo,
		Publisher: publisher,
	}

	h := &handler.PaymentHandler{
		Service: svc,
	}

	router := gin.Default()

	routes.RegisterRoutes(router, h)

	router.Run(":8082")
}