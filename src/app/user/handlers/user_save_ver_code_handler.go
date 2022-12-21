package handlers

import (
	"app/user/models"
	"app/user/repositories"
)

func UserSaveVerificationCodeHandler(user models.UserModel, code string) (models.UserModel, error) {
	userRepo := repositories.NewUserRepository()

	user, err := userRepo.UpdateVerificationCode(user, code)

	return user, err
}
