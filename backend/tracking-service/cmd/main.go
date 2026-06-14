package main

import (
	"tracking-service/internal/config"
	"tracking-service/internal/handler"
	"tracking-service/internal/repository"
	"tracking-service/internal/routes"
	"tracking-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

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