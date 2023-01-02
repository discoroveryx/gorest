package handlers

import (
	"app/user/models"

	"golang.org/x/crypto/bcrypt"
)

// 'user.Password' is a hash of password,
// 'password' is a clear password
func ValidatePasswordByUserHandler(user models.UserModel, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
