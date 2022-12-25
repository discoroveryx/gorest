package transport

import (
	auth_controller "app/auth/controllers"
	user_controller "app/user/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(debug bool) *gin.Engine {
	// fmt.Println("I am the SetupRouter")

	// router := gin.New()
	// router.Use(CORSMiddlware())

	if debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(cors.Default())

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}
	// router.Use(cors.New(config))

	router.Use(AuthMiddleware())

	// Routers
	router.POST("/user/create/", user_controller.UserCreateController)
	router.POST("/user/login/", auth_controller.UserLoginController)
	router.GET("/user/profile/", user_controller.UserProfileController)
	router.POST("/user/verify/", user_controller.UserVerifyController)

	return router
}
