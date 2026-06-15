package routes

import (
	"tracking-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	trackingHandler *handler.TrackingHandler,
) {

	router.POST(
		"/tracking",
		trackingHandler.CreateTracking,
	)
}