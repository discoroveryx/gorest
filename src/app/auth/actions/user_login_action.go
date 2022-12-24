package actions

import (
	"app/auth/exceptions"
	auth_handlers "app/auth/handlers"
	user_handlers "app/user/handlers"
	"fmt"

	// "app/auth/repositories"
	"app/auth/transformers"
	"app/user/models"
)

type UserLoginAction struct{}

func (a UserLoginAction) Run(serializerData transformers.UserLoginTransformer) (string, error) {
	var user models.UserModel

	// if !handlers.PasswordMinLengthValidator(&serializerData.Password) {
	// 	return user, exceptions.MinLengthInvalidPasswordError
	// 	// return user, &exceptions.MinLengthInvalidPasswordError{
	// 	// 	Title:  "password",
	// 	// 	Detail: "enter a valid password",
	// 	// 	// Err:    errors.New("enter_a_valid_password"),
	// 	// }
	// }

	// if !handlers.PasswordComparingValidator(&serializerData.Password, &serializerData.PasswordRepeated) {
	// 	return user, exceptions.PasswordComparingError
	// }

	// Check if user exists
	user, err := user_handlers.UserExistsByNameHandler(serializerData.Name)
	// fmt.Println("\n", user_exists, err, "\n")
	if err != nil {
		return "", exceptions.UserLoginFailedError
	}

	// Check is user verified
	userIsVerified, err := user_handlers.IsUserVerifiedByIdHandler(user.ID)
	fmt.Println("\n", userIsVerified, err, "\n")
	if err != nil {
		return "", exceptions.UserLoginFailedError
	}
	if userIsVerified == false {
		return "", exceptions.UserIsNotVerifiedError
	}

	// serializerData.Password, err = user_handlers.PasswordHashingHandler(serializerData.Password)
	// if err != nil {
	// 	fmt.Println("2")
	// 	return "", exceptions.UserLoginFailedError
	// }

	// fmt.Println(user.Password, serializerData.Password)
	// if user.Password != serializerData.Password {
	// 	return "", exceptions.UserLoginFailedError
	// }

	if !auth_handlers.ValidatePasswordByUserHandler(user, serializerData.Password) {
		fmt.Println("3")
		return "", exceptions.UserLoginFailedError
	}

	token, _ := auth_handlers.GenerateJWTByUserHandler(user.ID)

	fmt.Printf("Type %T", user.ID)

	fmt.Println("JWT = ", token)

	return token, nil
}
