package user

import (
	"app/user/models"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"tests/fixtures"
	"transport"
)

type AuthByJWTTestSuite struct {
	suite.Suite
	db     *gorm.DB
	router *gin.Engine
	fixtures.SuiteFixtures
}

func (suite *AuthByJWTTestSuite) SetupTest() {
	suite.db = suite.MockDatabase()
	suite.router = transport.SetupRouter(false)
}

func (suite *AuthByJWTTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&models.UserModel{})
}

func (suite *AuthByJWTTestSuite) TestAuth200() {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user/profile/", nil)

	user, _ := suite.CreateNewUserFixture(true)
	suite.PatchRequestWithJWT(*request, user.ID)

	suite.router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(200, response.StatusCode)
}

func (suite *AuthByJWTTestSuite) TestAuth401() {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user/profile/", nil)

	suite.router.ServeHTTP(recorder, request)

	response := recorder.Result()

	suite.Equal(401, response.StatusCode)
}

func TestRunnerAuthByJWT(t *testing.T) {
	suite.Run(t, new(AuthByJWTTestSuite))
}
