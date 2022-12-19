package user

import (
	"app/user/models"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"tests/fixtures"
	"transport"
)

type GetUserProfileTestSuite struct {
	suite.Suite
	db *gorm.DB
	fixtures.SuiteFixtures
}

func (suite *GetUserProfileTestSuite) SetupTest() {
	suite.db = suite.MockDatabase()
}

func (suite *GetUserProfileTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&models.UserModel{})
}

func (suite *GetUserProfileTestSuite) TestGetUserProfile200() {
	router := transport.SetupRouter()

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user/profile/", nil)

	user, _ := suite.CreateNewUserFixture()
	suite.PatchRequestWithJWT(*request, user.ID)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(200, response.StatusCode)

	// var responseData transformers.UserCreateRespTransformer

	// json.Unmarshal(recorder.Body.Bytes(), &responseData)

	// expectedData := transformers.UserCreateRespTransformer{
	// 	ID:        responseData.ID,
	// 	Name:      "vasya",
	// 	Email:     "vasya@vasya.com",
	// 	CreatedAt: responseData.CreatedAt,
	// }

	// suite.Equal(responseData, expectedData)
}

// func (suite *GetUserProfileTestSuite) TestUserCreate400() {
// 	handlers.CreateNewUserHandler{Repository: repositories.NewUserCreateRepository()}.Run(
// 		"vasya",
// 		"vasya@vasya.com",
// 		"12345678",
// 	)

// 	router := transport.SetupRouter()

// 	jsonBody := []byte(`{
// 		"name": "vasya",
// 		"email": "vasya@vasya.com",
// 		"password": "12345678",
// 		"password_repeated": "12345678"
// 	}`)
// 	bodyReader := bytes.NewReader(jsonBody)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/user/create/", bodyReader)
// 	router.ServeHTTP(w, req)

// 	suite.Equal(400, w.Code)
// 	// fmt.Println(w.Body.String())
// }

func TestRunnerUserGetProfile(t *testing.T) {
	suite.Run(t, new(GetUserProfileTestSuite))
}
