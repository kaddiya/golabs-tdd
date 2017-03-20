package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var signUpTests = []struct {
	Title      string
	Request    UserSignUpRequest
	Response   UserSignUpResponse
	StatusCode int
	Message    string
}{{
	Title:      "System should sign up a user and return a OK status when a unique email_id,proper password is supplied",
	Request:    UserSignUpRequest{UserName: "user1", Email: "user1@gmail.com", Password: "test1234"},
	Response:   UserSignUpResponse{UserName: "user1", UserID: 1},
	StatusCode: 400,
	Message:    "You have successfully signed up",
},
	{
		Title:      "System should throw a BAD_REQUEST when an invalid email_id is supplied",
		Request:    UserSignUpRequest{UserName: "user1", Email: "user1", Password: "test1234"},
		Response:   UserSignUpResponse{},
		StatusCode: 400,
		Message:    "Please enter a valid email id",
	},
	{
		Title:      "System should throw a CONFLICT error when an existing email is supplied",
		Request:    UserSignUpRequest{UserName: "user1", Email: "user1@gmail.com", Password: "test1234"},
		Response:   UserSignUpResponse{},
		StatusCode: 400,
		Message:    "This email has already been taken",
	},
	{
		Title:      "System should throw a BAD_REQUEST when a password with less than 8 characters in supplied",
		Request:    UserSignUpRequest{UserName: "user1", Email: "user3@gmail.com", Password: "test12"},
		Response:   UserSignUpResponse{},
		StatusCode: 400,
		Message:    "Password must be at least 8 characters long",
	},
}

func TestUserSignup(t *testing.T) {
	server := httptest.NewServer(InitRouter())
	for _, fixture := range signUpTests {
		t.Log("\n")
		t.Log("Executing test", fixture.Title)

		body, _ := json.Marshal(fixture.Request)

		reader := strings.NewReader(string(body))
		request, _ := http.NewRequest("POST", "/users/new", reader)
		w := httptest.NewRecorder()
		SignUpUser(w, request)
		if w.Code != fixture.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, fixture.StatusCode)
			t.Fail()
			continue
		}

		response := &UserSignUpResponse{}
		json.Unmarshal(w.Body.Bytes(), response)
		t.Log("The assertions for this test have passed")
	}

	server.Close()

}
