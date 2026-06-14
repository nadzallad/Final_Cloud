package main

import (
	"context"
	"log"

	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/routes"
	"order-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	db, err := pgxpool.New(
		context.Background(),
		"postgres://postgres:postgres@localhost:5432/logistic",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := gin.Default()

	router.Use(cors.Default())

	orderRepo := repository.NewOrderRepository(
		db,
	)

	orderService := service.NewOrderService(
		orderRepo,
	)

	orderHandler := handler.NewOrderHandler(
		orderService,
	)

	routes.SetupRoutes(
		router,
		orderHandler,
	)

	router.Run(":8081")
}