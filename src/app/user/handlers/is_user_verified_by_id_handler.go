package handlers

import (
	"app/user/repositories"
)

func IsUserVerifiedByIdHandler(id uint) (bool, error) {
	userRepo := repositories.NewUserRepository()

	return userRepo.IsUserVerifiedById(id)
}
