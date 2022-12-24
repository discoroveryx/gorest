package handlers

import (
	"app/user/repositories"
)

func UserExistsByEmailHandler(email string) bool {
	userRepo := repositories.NewUserRepository()

	return userRepo.UserExistsByEmail(email)
}
