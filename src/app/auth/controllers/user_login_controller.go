package controllers

import (
	"net/http"

	"app/auth/actions"
	"app/auth/transformers"

	"github.com/gin-gonic/gin"
)

func UserLoginController(c *gin.Context) {
	var serializerData transformers.UserLoginTransformer

	if err := c.ShouldBindJSON(&serializerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := actions.UserLoginAction{}.Run(serializerData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusCreated, gin.H{"status": "OK", "user_id": user.ID, "ctime": user.Ctime})

	resp := transformers.UserLoginRespTransformer{
		// UserID: result.UserID,
		// Token:  result.Token,
		Token: token,
	}
	c.JSON(http.StatusCreated, resp)
}
