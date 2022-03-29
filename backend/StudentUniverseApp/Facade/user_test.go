package server

import (
	database "StudentUniverse/StudentUniverseApp/Facade/Database"
	"bytes"
	"encoding/json"

	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SignUp_NewUser(t *testing.T) {

	testUsers := &database.Users{
		Email:    "testing1212@tester.com",
		Password: "tester123",
	}
	fmt.Print("New User Test - Unique user")
	userformValue, _ := json.Marshal(testUsers)
	r := SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/signup", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Test failed - Check User exists")

}
func Test_SignUp_UserAlreadyExists(t *testing.T) {
	testUsers := &database.Users{
		Email:    "testing2@tester.com",
		Password: "tester123",
	}
	userformValue, _ := json.Marshal(testUsers)
	fmt.Print("New User Test - Existing Users")
	//gin.SetMode(gin.TestMode)
	r := SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/signup", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Test failed - Check User exists")

}
func Test_SignUp_FormValidation(t *testing.T) {
	testUsers := &database.Users{
		Email:    "testing2@tester.com",
		Password: "123",
	}
	userformValue, _ := json.Marshal(testUsers)
	r := SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/signup", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Test failed - Form validation")
}
