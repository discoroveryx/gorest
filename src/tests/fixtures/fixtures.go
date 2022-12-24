package fixtures

import (
	auth_handlers "app/auth/handlers"
	"app/user/handlers"
	"app/user/models"
	"app/user/repositories"
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

	tokenKey, _ := auth_handlers.GenerateJWTByUserHandler(user.ID)
	fmt.Println("t", tokenKey)

	return user, tokenKey
}

func (s *SuiteFixtures) PatchRequestWithJWT(request http.Request, userId uint) {
	token, _ := auth_handlers.GenerateJWTByUserHandler(userId)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
}
