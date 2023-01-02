package fixtures

import (
	authhandlers "app/auth/handlers"
	"app/user/handlers"
	"app/user/models"
	"app/user/repositories"
	"app/user/transformers"
	"config"
	"dbstorage"
	"fmt"
	"helpers"

	"net/http"

	"gorm.io/gorm"
)

const DatabaseTestNameDefault = "test_main"
const DatabaseTestNamePrefix = "test_app_"

type SuiteFixtures struct{}

func (s *SuiteFixtures) MockDatabase() *gorm.DB {
	conf := config.GetProjectConf()

	conf.DataBase.Name = DatabaseTestNameDefault

	if catalogName, err := helpers.GetLastCatalogName(); err == nil {
		conf.DataBase.Name = fmt.Sprint(DatabaseTestNamePrefix, catalogName)
	}

	db := new(dbstorage.DB)
	cursor := db.Connect()
	db.Migrate()

	return cursor
}

func (s *SuiteFixtures) CreateNewUserFixture(verified bool) (models.UserModel, string) {
	serializerData := transformers.UserCreateTransformer{
		Name:             "vasya",
		Email:            "vasya@vasya.com",
		Password:         "12345678",
		PasswordRepeated: "12345678",
	}

	serializerData.Password, _ = handlers.PasswordHashingHandler(serializerData.Password)

	user := handlers.CreateNewUserHandler{Repository: repositories.NewUserCreateRepository()}.Run(
		serializerData.Name,
		serializerData.Email,
		serializerData.Password,
		verified,
	)

	tokenKey, _ := authhandlers.GenerateJWTByUserHandler(user.ID)

	return user, tokenKey
}

func (s *SuiteFixtures) PatchRequestWithJWT(request http.Request, userId uint) {
	token, _ := authhandlers.GenerateJWTByUserHandler(userId)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
}
