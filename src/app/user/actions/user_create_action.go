package actions

import (
	"app/user/exceptions"
	"app/user/handlers"
	"app/user/models"
	"app/user/repositories"
	"app/user/transformers"
)

// type UserCreateActioner interface {
// 	Run(transformers.UserCreateTransformer) (models.UserModel, error)
// }

type UserCreateAction struct{}

func (a UserCreateAction) Run(serializerData transformers.UserCreateTransformer) (models.UserModel, error) {
	var user models.UserModel

	if !handlers.PasswordMinLengthValidator(&serializerData.Password) {
		return user, exceptions.MinLengthInvalidPasswordError
		// return user, &exceptions.MinLengthInvalidPasswordError{
		// 	Title:  "password",
		// 	Detail: "enter a valid password",
		// 	// Err:    errors.New("enter_a_valid_password"),
		// }
	}

	if !handlers.PasswordComparingValidator(&serializerData.Password, &serializerData.PasswordRepeated) {
		return user, exceptions.PasswordComparingError
	}

	// Check if user exists
	if handlers.UserExistsByEmailHandler(serializerData.Email) {
		return user, exceptions.UserExistsByEmailError
	}

	_, err := handlers.UserExistsByNameHandler(serializerData.Name)
	if err == nil {
		return user, exceptions.UserExistsByNameError
	}

	serializerData.Password, _ = handlers.PasswordHashingHandler(serializerData.Password)

	user = handlers.CreateNewUserHandler{Repository: repositories.NewUserCreateRepository()}.Run(
		serializerData.Name,
		serializerData.Email,
		serializerData.Password,
		false,
	)

	verificationCode := handlers.GenerateVerificationCodeHandler()

	handlers.UserSaveVerificationCodeHandler(user, verificationCode)

	handlers.SendVerificationCodeByEmailHandler(verificationCode, user.Email)

	return user, nil
}
