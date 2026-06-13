package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/api/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"total_orders":   0,
			"total_payments": 0,
			"total_revenue":  0,
		})
	})

	router.Run(":8080")
}