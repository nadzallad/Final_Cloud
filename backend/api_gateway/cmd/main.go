package main

import (
	"net/http"

	"api_gateway/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	// ROUTE LAMA TETAP ADA
	router.GET("/api/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"total_orders":   0,
			"total_payments": 0,
			"total_revenue":  0,
		})
	})

	// ROUTE BARU UNTUK TEST JWT
	router.GET(
		"/api/me",
		middleware.AuthMiddleware(),
		func(c *gin.Context) {

			userID, _ := c.Get("userID")
			role, _ := c.Get("role")
			email, _ := c.Get("email")

			c.JSON(http.StatusOK, gin.H{
				"userID": userID,
				"role":   role,
				"email":  email,
			})
		},
	)
	router.GET(
		"/api/admin",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("admin"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome Admin",
			})
		},
	)

	router.GET(
		"/api/courier",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("kurir"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome Courier",
			})
		},
	)

	router.GET(
		"/api/user",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("user"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome User",
			})
		},
	)

	router.Run(":8080")
}
