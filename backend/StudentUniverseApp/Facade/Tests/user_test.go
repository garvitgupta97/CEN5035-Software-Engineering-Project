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
		Email:    "testemailMain323423@tester.com",
		Password: "tester1234",
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
	assert.Equal(t, http.StatusBadRequest, w.Code, "Test passed - Form validation")
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

func Test_Post_CreatePost(t *testing.T) {
	testPost := &database.Post{
		UserId:        1,
		ThreadId:      1,
		Title:         "title",
		Content:       "contnr",
		Votes:         2,
		CommentsCount: 0,
		ThreadTitle:   "thread",
	}
	userformValue, _ := json.Marshal(testPost)
	r := rtr.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/post/create", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Test passed - Creation")
}

func Test_Post_GetAllPosts(t *testing.T) {
	testPost := database.GetAllPosts()
	userformValue, _ := json.Marshal(testPost)
	r := rtr.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/post/allPosts", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Test passed - Get all posts")
}

func Test_Post_GetPostbyID(t *testing.T) {
	testPost := database.GetPostById(1)
	userformValue, _ := json.Marshal(testPost)
	r := rtr.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/post/1", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Test passed - Get post by id")
}
func Test_Comment_CreateComment(t *testing.T) {
	testPost := &database.Comment{
		CommentId:       1,
		ParentCommentId: 1,
		UserId:          1,
		PostId:          1,

		Content: "contnr",
	}
	userformValue, _ := json.Marshal(testPost)
	r := rtr.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/comment/create", bytes.NewBuffer(userformValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Test passed - Creation")
}
