package actions

import (
	"app/auth/exceptions"
	authhandlers "app/auth/handlers"
	userhandlers "app/user/handlers"

	// "app/auth/repositories"
	"app/auth/transformers"
	"app/user/models"
)

type UserLoginAction struct{}

func (a UserLoginAction) Run(serializerData transformers.UserLoginTransformer) (string, error) {
	var user models.UserModel

	// Check if user exists
	user, err := userhandlers.UserExistsByNameHandler(serializerData.Name)
	if err != nil {
		return "", exceptions.UserLoginFailedError
	}

	// Check is user verified
	userIsVerified, err := userhandlers.IsUserVerifiedByIdHandler(user.ID)
	if err != nil {
		return "", exceptions.UserLoginFailedError
	}
	if userIsVerified == false {
		return "", exceptions.UserIsNotVerifiedError
	}

	if !authhandlers.ValidatePasswordByUserHandler(user, serializerData.Password) {
		return "", exceptions.UserLoginFailedError
	}

	token, _ := authhandlers.GenerateJWTByUserHandler(user.ID)

	return token, nil
}
