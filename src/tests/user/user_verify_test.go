package user

import (
	"app/user/models"
	"app/user/transformers"
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

type UserVerifyTestSuite struct {
	suite.Suite
	db *gorm.DB
	fixtures.SuiteFixtures
}

func (suite *UserVerifyTestSuite) SetupTest() {
	suite.db = suite.MockDatabase()
}

func (suite *UserVerifyTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&models.UserModel{})
}

func (suite *UserVerifyTestSuite) TestUserVerify200() {
	user, token := suite.CreateNewUserFixture(true)

	fmt.Println("token on tests", user.VerificationCode)

	router := transport.SetupRouter()

	userVerifyData := &transformers.UserVerifyTransformer{
		UserId:           user.ID,
		VerificationCode: token,
	}

	jsonBody, _ := json.Marshal(userVerifyData)
	bodyReader := bytes.NewReader(jsonBody)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/user/verify/", bodyReader)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	fmt.Println("11", response)

	suite.Equal(201, response.StatusCode)

	var responseData transformers.UserVerifyRespTransformer

	json.Unmarshal(recorder.Body.Bytes(), &responseData)

	fmt.Println(user.ID, user.Verified)
	fmt.Println(responseData)

	expectedData := transformers.UserVerifyRespTransformer{UserId: user.ID, Verified: user.Verified}

	suite.Equal(&responseData, &expectedData)
}

func (suite *UserVerifyTestSuite) TestUserNotVerified400() {}

func TestRunnerUserVerify(t *testing.T) {
	suite.Run(t, new(UserVerifyTestSuite))
}
