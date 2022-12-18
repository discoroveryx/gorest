package user

import (
	"app/user/handlers"
	"app/user/models"
	"app/user/repositories"
	"app/user/transformers"
	"bytes"
	"config"
	"dbstorage"
	"encoding/json"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"transport"
)

type CreateUserTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *CreateUserTestSuite) SetupTest() {
	// suite.myAddExpected = 6

	// TODO Move it to TestSuit, to follow to DRY
	conf := config.GetProjectConf()
	conf.DBName = "test_1.db"

	db := new(dbstorage.DB)
	suite.db = db.Connect()
	db.Migrate()
}

func (suite *CreateUserTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&models.UserModel{})
}

func (suite *CreateUserTestSuite) TestUserCreate400() {
	handlers.CreateNewUserHandler{Repository: repositories.NewUserCreateRepository()}.Run(
		"vasya",
		"vasya@vasya.com",
		"12345678",
	)

	router := transport.SetupRouter()

	jsonBody := []byte(`{
		"name": "vasya",
		"email": "vasya@vasya.com",
		"password": "12345678",
		"password_repeated": "12345678"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/create/", bodyReader)
	router.ServeHTTP(w, req)

	suite.Equal(400, w.Code)
	// fmt.Println(w.Body.String())
}

func (suite *CreateUserTestSuite) TestUserCreate201() {
	router := transport.SetupRouter()

	newUser := &transformers.UserCreateTransformer{
		Name:             "vasya",
		Email:            "vasya@vasya.com",
		Password:         "12345678",
		PasswordRepeated: "12345678",
	}

	jsonBody, _ := json.Marshal(newUser)

	bodyReader := bytes.NewReader(jsonBody)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/user/create/", bodyReader)
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(201, response.StatusCode)

	var responseData transformers.UserCreateRespTransformer

	json.Unmarshal(recorder.Body.Bytes(), &responseData)

	expectedData := transformers.UserCreateRespTransformer{
		ID:        responseData.ID,
		Name:      "vasya",
		Email:     "vasya@vasya.com",
		CreatedAt: responseData.CreatedAt,
	}

	suite.Equal(responseData, expectedData)
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(CreateUserTestSuite))
}
