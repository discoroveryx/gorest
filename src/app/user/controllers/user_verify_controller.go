package controllers

import (
	"net/http"

	"app/user/actions"
	"app/user/transformers"

	"github.com/gin-gonic/gin"
)

func UserVerifyController(c *gin.Context) {
	var serializerData transformers.UserVerifyTransformer

	if err := c.ShouldBindJSON(&serializerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := actions.UserVerifyAction{}.Run(serializerData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := transformers.UserVerifyRespTransformer{
		UserId:   user.ID,
		Verified: true,
	}
	c.JSON(http.StatusOK, resp)
}
