package user

import (
	"app/dbstorage"
	"bytes"
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"

	"transport"
)

type CreateUserTestSuite struct {
	suite.Suite
	// myAddExpected int
}

func (suite *CreateUserTestSuite) SetupTest() {
	// suite.myAddExpected = 6
	db := dbstorage.NewTestDB()
	db.Migrate()
}

func (suite *CreateUserTestSuite) TestCreateUser() {
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
	fmt.Println(w.Body.String())
	// suite.Equal("pong", w.Body.String())
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(CreateUserTestSuite))
}
