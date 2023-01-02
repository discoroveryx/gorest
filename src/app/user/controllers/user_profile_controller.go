package controllers

import (
	"net/http"

	"app/user/actions"
	"app/user/transformers"

	"github.com/gin-gonic/gin"
)

func UserProfileController(c *gin.Context) {
	if c.GetBool("Authenticated") == false {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := c.GetUint("UserId")

	user, err := actions.GetUserProfileAction{}.Run(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	serializerData := transformers.UserProfileTransformer{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(200, serializerData)
}
