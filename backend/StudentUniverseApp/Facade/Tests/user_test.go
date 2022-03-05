package tests

import (
	database "StudentUniverse/StudentUniverseApp/Facade/Database"
	"bytes"
	"encoding/json"

	rtr "StudentUniverse/StudentUniverseApp/Facade"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SignUp_NewUser(t *testing.T) {

	testUsers := &database.Users{
		Email:    "testing2@tester.com",
		Password: "tester123",
	}
	fmt.Print("New User Test - Unique user")
	userformValue, _ := json.Marshal(testUsers)
	r := rtr.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/signup", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Test passed - Add new users")

}
func Test_SignUp_UserAlreadyExists(t *testing.T) {

	testUsers := &database.Users{
		Email:    "testing2@tester.com",
		Password: "tester123",
	}
	userformValue, _ := json.Marshal(testUsers)
	fmt.Print("New User Test - Existing Users")
	//gin.SetMode(gin.TestMode)
	r := rtr.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/signup", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Test passed - Check User exists")
}
func Test_SignUp_FormValidation(t *testing.T) {
	testUsers := &database.Users{
		Email:    "testing2@tester.com",
		Password: "123",
	}
	userformValue, _ := json.Marshal(testUsers)
	r := rtr.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/signup", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Test passed - Form validation")
}

func Test_SignIn_FormValidation(t *testing.T) {
	testUsers := &database.Users{
		Email:    "testing2@tester.com",
		Password: "123",
	}
	userformValue, _ := json.Marshal(testUsers)
	r := rtr.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/signin", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Test passed - Form validation")
}

func Test_SignIn_SinginSuccess(t *testing.T) {
	testUsers := &database.Users{
		Email:    "testing2@tester.com",
		Password: "tester123",
	}
	userformValue, _ := json.Marshal(testUsers)
	r := rtr.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/signin", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Test passed - Login")
}
