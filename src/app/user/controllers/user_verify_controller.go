package controllers

import (
	"fmt"
	"net/http"

	"app/user/actions"
	"app/user/transformers"

	"github.com/gin-gonic/gin"
)

func UserVerifyController(c *gin.Context) {
	var serializerData transformers.UserVerifyTransformer

	fmt.Println("UserVerifyController")
	if err := c.ShouldBindJSON(&serializerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	fmt.Println("ser data", serializerData)

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
