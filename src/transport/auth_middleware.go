package transport

import (
	"strings"

	auth_handlers "app/auth/handlers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// fmt.Println("authHeader", authHeader)

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

		// fmt.Println("bearer", bearer)

		claims, err := auth_handlers.ParseJWTHandler(bearer)

		if err != nil {
			return
		}
		// fmt.Println(err)
		// fmt.Println(claims)

		// fmt.Printf("Type %T", claims.UserId)

		c.Set("Authenticated", true)
		c.Set("UserId", claims.UserId)

		// if err != nil {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

	}
}
