package transport

import (
	"fmt"
	"strings"

	auth_controller "app/auth/controllers"
	auth_handlers "app/auth/handlers"
	user_controller "app/user/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,Content-Length,Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "43200")
	}
}

func AuthMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println("authHeader", authHeader)

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

		fmt.Println("bearer", bearer)

		claims, err := auth_handlers.ParseJWTHandler(bearer)

		if err != nil {
			return
		}
		fmt.Println(err)
		fmt.Println(claims)

		fmt.Printf("Type %T", claims.UserId)

		c.Set("Authenticated", true)
		c.Set("UserId", claims.UserId)

		// if err != nil {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

	}
}

func SetupRouter() *gin.Engine {
	fmt.Println("I am the SetupRouter")

	// router := gin.New()
	// router.Use(CORSMiddlware())

	router := gin.Default()
	router.Use(cors.Default())

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}
	// router.Use(cors.New(config))

	router.Use(AuthMiddlware())

	// Routers
	router.POST("/user/create/", user_controller.UserCreateController)
	router.POST("/user/login/", auth_controller.UserLoginController)
	router.GET("/user/profile/", user_controller.UserProfileController)

	return router
}
