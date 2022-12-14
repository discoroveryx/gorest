package handlers

import (
	"app/user/models"
	"app/user/repositories"
)

func GetUserByIdHandler(id uint) (models.UserModel, error) {
	user_repo := repositories.NewUserRepository()

	user, err := user_repo.GetUserById(id)

	return user, err
}
