package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("cloud_logistic_secret")

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token required",
			})

			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return JwtSecret, nil
			},
		)

		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})

			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("userID", claims["id"])
		c.Set("role", claims["role"])
		c.Set("email", claims["email"])

		c.Next()
	}
}