package user

import (
	"app/user/handlers"
	"app/user/models"
	"app/user/transformers"
	"bytes"
	"encoding/json"
	"testing"
	"tests/fixtures"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"transport"
)

type CreateUserTestSuite struct {
	suite.Suite
	db *gorm.DB
	fixtures.SuiteFixtures
}

func (suite *CreateUserTestSuite) SetupTest() {
	suite.db = suite.MockDatabase()
}

func (suite *CreateUserTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&models.UserModel{})
}

func (suite *CreateUserTestSuite) TestUserCreate400() {
	suite.CreateNewUserFixture()

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

	var responseData transformers.UserCreateResponseTransformer

	json.Unmarshal(recorder.Body.Bytes(), &responseData)

	expectedData := transformers.UserCreateResponseTransformer{
		ID:        responseData.ID,
		Name:      "vasya",
		Email:     "vasya@vasya.com",
		CreatedAt: responseData.CreatedAt,
		Verified:  false,
	}

	suite.Equal(&responseData, &expectedData)

	userVerified, _ := handlers.IsUserVerifiedByIdHandler(responseData.ID)
	suite.Equal(userVerified, false)
}

func TestRunnerUserCreate(t *testing.T) {
	suite.Run(t, new(CreateUserTestSuite))
}
