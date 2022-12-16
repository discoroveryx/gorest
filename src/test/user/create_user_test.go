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

	"github.com/golang/mock/gomock"
)

type CreateUserTestSuite struct {
	suite.Suite
	// myAddExpected int
}

func (suite *CreateUserTestSuite) SetupTest() {
	// suite.myAddExpected = 6
	// db := dbstorage.NewDB(new(dbstorage.MyMockDb))
	// db.Migrate()

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

	// suite.Equal(400, w.Code)
	fmt.Println(w.Body.String())
	// suite.Equal("pong", w.Body.String())
}

func TestRunner(t *testing.T) {
	// db := dbstorage.NewDB(new(dbstorage.MyMockDb))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := dbstorage.NewMockMyDbInter(ctrl)
	m.
		EXPECT().
		NewConn().
		Return().
		Times(0)

	suite.Run(t, new(CreateUserTestSuite))
}
