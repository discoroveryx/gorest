package user

import (
	"app/auth/transformers"
	"app/user/models"
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"tests/fixtures"
	"transport"
)

type UserLoginTestSuite struct {
	suite.Suite
	db *gorm.DB
	fixtures.SuiteFixtures
}

func (suite *UserLoginTestSuite) SetupTest() {
	suite.db = suite.MockDatabase()
}

func (suite *UserLoginTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&models.UserModel{})
}

func (suite *UserLoginTestSuite) TestUserLogin201() {
	suite.CreateNewUserFixture(true)

	router := transport.SetupRouter()

	userLoginData := &transformers.UserLoginTransformer{
		Name:     "vasya",
		Password: "12345678",
	}

	jsonBody, _ := json.Marshal(userLoginData)
	bodyReader := bytes.NewReader(jsonBody)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/user/login/", bodyReader)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(201, response.StatusCode)

	var responseData transformers.UserLoginRespTransformer

	json.Unmarshal(recorder.Body.Bytes(), &responseData)

	expectedData := transformers.UserLoginRespTransformer{Token: responseData.Token}

	suite.Equal(&responseData, &expectedData)
}
func (suite *UserLoginTestSuite) TestUserLoginNameWrong400() {
	suite.CreateNewUserFixture(true)

	router := transport.SetupRouter()

	userLoginData := &transformers.UserLoginTransformer{
		Name:     "wrong",
		Password: "12345678",
	}

	jsonBody, _ := json.Marshal(userLoginData)
	bodyReader := bytes.NewReader(jsonBody)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/user/login/", bodyReader)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(400, response.StatusCode)
	fmt.Println(recorder.Body.String())

	responseError := struct{ Error string }{}

	json.Unmarshal(recorder.Body.Bytes(), &responseError)

	expectedError := struct{ Error string }{Error: "user_login_failed"}

	suite.Equal(&responseError, &expectedError)
}

func (suite *UserLoginTestSuite) TestUserLoginNamePassword400() {
	suite.CreateNewUserFixture(true)

	router := transport.SetupRouter()

	userLoginData := &transformers.UserLoginTransformer{
		Name:     "vasya",
		Password: "wrong_password",
	}

	jsonBody, _ := json.Marshal(userLoginData)
	bodyReader := bytes.NewReader(jsonBody)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/user/login/", bodyReader)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(400, response.StatusCode)
	fmt.Println(recorder.Body.String())

	responseError := struct{ Error string }{}

	json.Unmarshal(recorder.Body.Bytes(), &responseError)

	expectedError := struct{ Error string }{Error: "user_login_failed"}

	suite.Equal(&responseError, &expectedError)
}

func (suite *UserLoginTestSuite) TestUserLoginNotVerified400() {
	suite.CreateNewUserFixture(false)

	router := transport.SetupRouter()

	userLoginData := &transformers.UserLoginTransformer{
		Name:     "vasya",
		Password: "12345678",
	}

	jsonBody, _ := json.Marshal(userLoginData)
	bodyReader := bytes.NewReader(jsonBody)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/user/login/", bodyReader)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(400, response.StatusCode)
	fmt.Println(recorder.Body.String())

	responseError := struct{ Error string }{}

	json.Unmarshal(recorder.Body.Bytes(), &responseError)

	expectedError := struct{ Error string }{Error: "user_is_not_verified"}

	suite.Equal(&responseError, &expectedError)
}

func TestRunnerUserLogin(t *testing.T) {
	suite.Run(t, new(UserLoginTestSuite))
}
