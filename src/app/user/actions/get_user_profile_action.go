package actions

import (
	"app/user/handlers"
	"app/user/models"
)

type GetUserProfileAction struct{}

func (a GetUserProfileAction) Run(userId uint) (models.UserModel, error) {
	var user models.UserModel

	user, err := handlers.GetUserByIdHandler(userId)

	return user, err
}
