package handlers

import "fmt"

func SendVerificationCodeByEmailHandler(code string, email string) bool {
	// TODO Turn into logger
	fmt.Printf("\nVerificationCode '%s' was sended to %s \n\n", code, email)

	return true
}
