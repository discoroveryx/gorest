package actions

import (
	"app/user/exceptions"
	"app/user/handlers"
	"app/user/models"
	"app/user/transformers"
)

type UserVerifyAction struct{}

func (a UserVerifyAction) Run(serializerData transformers.UserVerifyTransformer) (models.UserModel, error) {
	var user models.UserModel
	var err error

	// // Check if user exists
	// if handlers.UserExistsByEmailHandler(serializerData.Email) {
	// 	return user, exceptions.UserExistsByEmailError
	// }

	user, err = handlers.UserVerifyHandler(serializerData.UserId, serializerData.VerificationCode)
	// fmt.Println("\n", user_exists, err, "\n")
	if err != nil {
		return user, exceptions.UserVerifyError
	}

	return user, nil
}
