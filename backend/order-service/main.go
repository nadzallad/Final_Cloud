package main

import (
	"log"
	"order-service/internal/config"
	"order-service/internal/entity"
	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/routes"
	"order-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	log.Println("DB connected!")

	db.AutoMigrate(&entity.Order{})

	orderRepo := repository.NewOrderRepository(db)
	cityRepo := repository.NewCityRepository(db)
	orderService := service.NewOrderService(orderRepo, cityRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router, orderHandler)

	log.Println("Order Service running on :8081")
	router.Run(":8081")
}