package controllers

import (
	"fmt"
	"net/http"

	"app/user/handlers"

	"app/user/transformers"

	"github.com/gin-gonic/gin"
)

func UserProfileController(c *gin.Context) {
	// var user models.UserModel

	if c.GetBool("Authenticated") == false {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := c.GetUint("UserId")

	// Add actions fo this handler
	user, _ := handlers.GetUserByIdHandler(userId)

	serializerData := transformers.UserCProfileTransformer{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	fmt.Println("\n\nuser", user)

	c.JSON(200, serializerData)
}
