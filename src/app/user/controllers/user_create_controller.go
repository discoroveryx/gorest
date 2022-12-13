package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"app/user/actions"
	"app/user/exceptions"
	"app/user/transformers"

	"github.com/gin-gonic/gin"
)

func UserCreateController(c *gin.Context) {
	fmt.Println("create user")
	var serializerData transformers.UserCreateTransformer

	if err := c.ShouldBindJSON(&serializerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := actions.UserCreateAction{}.Run(serializerData)
	if err != nil {
		switch {
		case errors.Is(err, exceptions.MinLengthInvalidPasswordError), errors.Is(err, exceptions.PasswordComparingError):
			c.JSON(http.StatusBadRequest, gin.H{"password": err.Error()})
		case errors.Is(err, exceptions.UserExistsByNameError):
			c.JSON(http.StatusBadRequest, gin.H{"name": fmt.Sprintf("user with name='%s' exists", serializerData.Name)})
		case errors.Is(err, exceptions.UserExistsByEmailError):
			c.JSON(http.StatusBadRequest, gin.H{"name": fmt.Sprintf("user with email='%s' exists", serializerData.Email)})
		}
		return
	}

	// c.JSON(http.StatusCreated, gin.H{"status": "OK", "user_id": user.ID, "ctime": user.Ctime})

	resp := transformers.UserCreateRespTransformer{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	c.JSON(http.StatusCreated, resp)
}
