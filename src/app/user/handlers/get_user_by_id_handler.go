package handlers

import (
	"app/user/models"
	"app/user/repositories"
)

func GetUserByIdHandler(id uint) (models.UserModel, error) {
	userRepo := repositories.NewUserRepository()

	user, err := userRepo.GetUserById(id)

	return user, err
}
