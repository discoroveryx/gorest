package handlers

import (
	"app/user/models"
	"app/user/repositories"
)

type CreateNewUserHandler struct {
	Repository repositories.UserCreateRepository
}

func (s CreateNewUserHandler) Run(name string, email string, password string, verified bool) models.UserModel {
	userRepo := s.Repository

	return userRepo.UserCreate(name, email, password, verified)
}
