package main

import (
	"app/dbstorage"

	auth_controller "app/auth/controllers"
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

func main() {
	// Init Database
	db := dbstorage.NewDB()
	db.Migrate()

	// Init http
	// router := gin.New()
	// router.Use(CORSMiddlware())

	router := gin.Default()
	router.Use(cors.Default())

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}
	// router.Use(cors.New(config))

	// Routers
	router.GET("/user/profile/", user_controller.UserDetailController)
	router.POST("/user/create/", user_controller.UserCreateController)
	router.POST("/user/login/", auth_controller.UserLoginController)

	router.Run()
}
