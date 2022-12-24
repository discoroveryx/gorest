package handlers

import (
	"app/user/models"
	"app/user/repositories"
)

func UserExistsByNameHandler(name string) (models.UserModel, error) {
	userRepo := repositories.NewUserRepository()

	user, err := userRepo.UserExistsByName(name)

	return user, err
}
