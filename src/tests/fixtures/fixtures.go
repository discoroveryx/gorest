package fixtures

import (
	auth_handlers "app/auth/handlers"
	"app/user/actions"
	"app/user/models"
	"app/user/transformers"
	"config"
	"dbstorage"
	"fmt"

	"net/http"

	"gorm.io/gorm"
)

const DatabaseTestName = "test_1"

type SuiteFixtures struct {
}

func (s *SuiteFixtures) MockDatabase() *gorm.DB {
	conf := config.GetProjectConf()
	conf.DBName = DatabaseTestName

	db := new(dbstorage.DB)
	cursor := db.Connect()
	db.Migrate()

	return cursor
}

func (s *SuiteFixtures) CreateNewUserFixture() (models.UserModel, error) {
	serializerData := transformers.UserCreateTransformer{
		Name:             "vasya",
		Email:            "vasya@vasya.com",
		Password:         "12345678",
		PasswordRepeated: "12345678",
	}

	user, err := actions.UserCreateAction{}.Run(serializerData)

	return user, err
}

func (s *SuiteFixtures) PatchRequestWithJWT(request http.Request, userId uint) {
	token, _ := auth_handlers.GenerateJWTByUserHandler(userId)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
}
