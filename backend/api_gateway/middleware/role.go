package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {

		role, exists := c.Get("role")

		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Role not found",
			})
			c.Abort()
			return
		}

		userRole := role.(string)

		for _, allowed := range roles {
			if userRole == allowed {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"message": "Access denied",
		})
		c.Abort()
	}
}