package handlers

import (
	"app/user/models"
	"app/user/repositories"
)

func UserVerifyHandler(userId uint, verificationCode string) (models.UserModel, error) {
	userRepo := repositories.NewUserRepository()

	user, err := userRepo.UserVerify(userId, verificationCode)

	return user, err
}
