package user

import (
	"app/user/models"
	"app/user/transformers"
	"encoding/json"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"tests/fixtures"
	"transport"
)

type GetUserProfileTestSuite struct {
	suite.Suite
	db     *gorm.DB
	router *gin.Engine
	fixtures.SuiteFixtures
}

func (suite *GetUserProfileTestSuite) SetupTest() {
	suite.db = suite.MockDatabase()
	suite.router = transport.SetupRouter(false)
}

func (suite *GetUserProfileTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&models.UserModel{})
}

func (suite *GetUserProfileTestSuite) TestGetUserProfile200() {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user/profile/", nil)

	user, _ := suite.CreateNewUserFixture(true)
	suite.PatchRequestWithJWT(*request, user.ID)

	suite.router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(200, response.StatusCode)

	var responseData transformers.UserProfileTransformer

	json.Unmarshal(recorder.Body.Bytes(), &responseData)

	expectedData := transformers.UserProfileTransformer{
		ID:        responseData.ID,
		Name:      "vasya",
		Email:     "vasya@vasya.com",
		CreatedAt: responseData.CreatedAt,
	}

	suite.Equal(&responseData, &expectedData)
}

func (suite *GetUserProfileTestSuite) TestGetUserProfile401() {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user/profile/", nil)

	suite.router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(401, response.StatusCode)
}

func TestRunnerUserGetProfile(t *testing.T) {
	suite.Run(t, new(GetUserProfileTestSuite))
}
