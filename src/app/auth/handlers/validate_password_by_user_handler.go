package handlers

import (
	"app/user/models"

	"golang.org/x/crypto/bcrypt"
)

// 'user.Password' is a hash of password,
// 'password' is a clear password
func ValidatePasswordByUserHandler(user models.UserModel, password string) bool {
	// fmt.Println(user.Password, password)
	// fmt.Println([]byte(user.Password), []byte(password))

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
