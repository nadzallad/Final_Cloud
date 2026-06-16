package routes

import (
	"warehouse-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, warehouseHandler *handler.WarehouseHandler) {
	api := router.Group("/api")
	{
		api.POST("/warehouse-logs", warehouseHandler.CreateLog)
		api.GET("/warehouse-logs", warehouseHandler.GetLogs)
		api.GET("/warehouse-logs/:id", warehouseHandler.GetLogByID)
		api.PATCH("/warehouse-logs/:id/status", warehouseHandler.UpdateLogStatus)
	}
}
