package mytests

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

type CreateUserTestSuite struct {
	suite.Suite
	myAddExpected int
}

func (suite *CreateUserTestSuite) SetupTest() {
	suite.myAddExpected = 6
}

func (suite *CreateUserTestSuite) TestUserCreate() {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/profile/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(CreateUserTestSuite))
}
