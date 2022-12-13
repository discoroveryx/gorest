package main

import (
	"app/dbstorage"

	// auth_controller "app/auth/controllers"
	user_controller "app/user/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init Database
	db := dbstorage.NewDB()
	db.Migrate()

	// Init http
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/user/profile/", user_controller.UserDetailController)
	router.POST("/user/create/", user_controller.UserCreateController)
	// router.POST("/user/login/", auth_controller.UserLoginController)

	router.Run()
}
