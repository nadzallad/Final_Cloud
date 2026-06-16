package main

import (
	"time"

	"tracking-service/internal/config"
	"tracking-service/internal/handler"
	"tracking-service/internal/repository"
	"tracking-service/internal/routes"
	"tracking-service/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	trackingRepo :=
		repository.NewTrackingRepository(
			config.TrackingCollection,
		)

	trackingService :=
		service.NewTrackingService(
			trackingRepo,
		)

	trackingHandler :=
		handler.NewTrackingHandler(
			trackingService,
		)

	routes.SetupRoutes(
		router,
		trackingHandler,
	)

	router.Run(":8087")
}