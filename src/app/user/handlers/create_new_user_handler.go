package handlers

import (
	"app/user/models"
	"app/user/repositories"
)


type CreateNewUserHandler struct {
	Repository repositories.UserCreateRepository
}

func (s CreateNewUserHandler) Run(name string, email string, password string) models.UserModel {
	user_repo := s.Repository

	return user_repo.UserCreate(name, email, password)
}
