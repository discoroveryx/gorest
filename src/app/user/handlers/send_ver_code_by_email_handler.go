package handlers

import "fmt"

func SendVerificationCodeByEmailHandler(code string, email string) bool {
	// user_repo := repositories.NewUserRepository()
	fmt.Printf("\nVerificationCode '%s' was sended to %s \n\n", code, email)

	return true
}
