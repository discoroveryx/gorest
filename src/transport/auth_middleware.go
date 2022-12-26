package transport

import (
	"strings"

	authhandlers "app/auth/handlers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) < 7 {
			return
		}

		valueSplited := strings.Split(authHeader, " ")

		if len(valueSplited) != 2 {
			return
		}

		if valueSplited[0] != "Bearer" {
			return
		}

		bearer := valueSplited[1]

		claims, err := authhandlers.ParseJWTHandler(bearer)

		if err != nil {
			return
		}

		c.Set("Authenticated", true)
		c.Set("UserId", claims.UserId)

	}
}
