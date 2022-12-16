package user

import (
	"app/dbstorage"
	"bytes"
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"transport"

	"github.com/golang/mock/gomock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser2(t *testing.T) {
	conn, _ := gorm.Open(sqlite.Open("test_2.db"), &gorm.Config{})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := dbstorage.NewMockMyDbInter(ctrl)
	m.
		EXPECT().
		NewConn().
		Return(conn).
		Times(1)

	// dbstorage.NewDB1(m)

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
	fmt.Println("w.Code", w.Code)
	fmt.Println("w.Body.String()", w.Body.String())
	// suite.Equal("pong", w.Body.String())

	// db := new(dbstorage.MyDb)
	dbstorage.NewDB1(m)
}
