package user

import (
	"app/user/models"
	"app/user/transformers"
	"bytes"
	"dbstorage"
	"encoding/json"
	"fmt"
	"myconfig"
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

	myconf := myconfig.GetMyConfig()
	dbname := myconf.DBName
	fmt.Println("\ndbname\n", dbname)

	myconf.DBName = "test_1.db"

	db := new(dbstorage.DB)
	cursor := db.Connect()

	cursor.Migrator().DropTable(&models.UserModel{})

	db.Migrate()
}

func (suite *CreateUserTestSuite) TestUserCreate400() {
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
	// suite.Equal("pong", w.Body.String())
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
