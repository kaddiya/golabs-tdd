package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var signUpTests = []struct {
	Title         string
	Request       UserSignUpRequest
	Response      UserSignUpResponse
	StatusCode    int
	Message       string
	URLPath       string
	mockedUserDao MockedUserDao
}{{
	Title:      "System should sign up a user and return a OK status when a unique email_id,proper password is supplied",
	Request:    UserSignUpRequest{UserName: "user1", Email: "user1@gmail.com", Password: "test1234"},
	Response:   UserSignUpResponse{UserName: "user1", UserID: 1},
	StatusCode: 200,
	Message:    "You have successfully signed up",
	URLPath:    "/users/new",
	mockedUserDao: MockedUserDao{
		isEmailIDUniqueFunc: func(email string) (bool, error) {
			return false, nil
		},
		saveUserFunc: func(u *UserSignUpRequest) {

		},
	},
},
	{
		Title:      "System should throw a BAD_REQUEST when an invalid email_id is supplied",
		Request:    UserSignUpRequest{UserName: "user1", Email: "user1", Password: "test1234"},
		Response:   UserSignUpResponse{},
		StatusCode: 400,
		Message:    "Please enter a valid email id",
		URLPath:    "/users/new",
		mockedUserDao: MockedUserDao{
			isEmailIDUniqueFunc: func(email string) (bool, error) {
				return false, nil
			},
			saveUserFunc: func(u *UserSignUpRequest) {

			},
		},
	},
	{
		Title:      "System should throw a CONFLICT error when an existing email is supplied",
		Request:    UserSignUpRequest{UserName: "user1", Email: "user1@gmail.com", Password: "test1234"},
		Response:   UserSignUpResponse{},
		StatusCode: 409,
		Message:    "This email has already been taken",
		URLPath:    "/users/new",
		mockedUserDao: MockedUserDao{
			isEmailIDUniqueFunc: func(email string) (bool, error) {
				return false, errors.New("This email has already been taken")
			},
			saveUserFunc: func(u *UserSignUpRequest) {

			},
		},
	},
	{
		Title:      "System should throw a BAD_REQUEST when a password with less than 8 characters in supplied",
		Request:    UserSignUpRequest{UserName: "user1", Email: "user3@gmail.com", Password: "test12"},
		Response:   UserSignUpResponse{},
		StatusCode: 400,
		Message:    "Password must be at least 8 characters long",
		URLPath:    "/users/new",
		mockedUserDao: MockedUserDao{
			isEmailIDUniqueFunc: func(email string) (bool, error) {
				return false, nil
			},
			saveUserFunc: func(u *UserSignUpRequest) {

			},
		},
	},
}

func TestUserSignup(t *testing.T) {
	c := &Container{Dao: &MockedUserDao{}}
	r := c.InitRouter()
	server := httptest.NewServer(r)
	for _, fixture := range signUpTests {
		t.Log("\n")
		t.Log("Executing test", fixture.Title)

		body, _ := json.Marshal(fixture.Request)

		reader := strings.NewReader(string(body))
		request, _ := http.NewRequest("POST", fixture.URLPath, reader)
		var match mux.RouteMatch
		b := r.Match(request, &match)
		if !b {
			t.Logf("could not find the route %s", fixture.URLPath)
			t.Fail()
		}
		w := httptest.NewRecorder()
		t.Log(fixture.Title)
		c.Dao = &fixture.mockedUserDao
		handlerFunc := http.HandlerFunc(c.SignUpUser)
		handlerFunc.ServeHTTP(w, request)
		//validate the API codes
		if w.Code != fixture.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, fixture.StatusCode)
			t.Fail()
			continue
		}
		//validate the error messages
		if string(w.Body.Bytes()) != fixture.Message {
			t.Logf("expected message to be %s but got %s", fixture.Message, string(w.Body.Bytes()))
			t.Fail()
			continue
		}

		response := &UserSignUpResponse{}
		json.Unmarshal(w.Body.Bytes(), response)
		t.Log("The assertions for this test have passed")
	}

	server.Close()

}

type MockedUserDao struct {
	// isEmailIDUniqueFuncFunc mocks the isEmailIDUniqueFunc function.
	isEmailIDUniqueFunc func(email string) (bool, error)
	// saveUserFuncFunc mocks the saveUserFunc function.
	saveUserFunc func(u *UserSignUpRequest)
}

func (mock MockedUserDao) isEmailIDUnique(email string) (bool, error) {
	if mock.isEmailIDUniqueFunc == nil {
		panic("moq: UserDaoMock.isEmailIDUniqueFunc is nil but was just called")
	}

	return mock.isEmailIDUniqueFunc(email)

}

// saveUser calls saveUserFunc.
func (mock *MockedUserDao) saveUser(u *UserSignUpRequest) {
	if mock.saveUserFunc == nil {
		panic("moq: UserDaoMock.saveUserFunc is nil but was just called")
	}

	mock.saveUserFunc(u)

}
