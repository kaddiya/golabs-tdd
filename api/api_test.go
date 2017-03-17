package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var signUpTests = []struct {
	title      string
	request    UserSignUpRequest
	response   UserSignUpResponse
	statusCode int
	message    string
}{{
	title:      "System should sign up a user and return a OK status when a unique email_id,proper password is supplied",
	request:    UserSignUpRequest{userName: "user1", email: "user1@gmail.com", password: "test1234"},
	response:   UserSignUpResponse{userName: "user1", userID: 1},
	statusCode: 200,
	message:    "You have successfully signed up",
},
	{
		title:      "System should throw a BAD_REQUEST when an invalid email_id is supplied",
		request:    UserSignUpRequest{userName: "user1", email: "user1", password: "test1234"},
		response:   UserSignUpResponse{},
		statusCode: 400,
		message:    "Please enter a valid email id",
	},
	{
		title:      "System should throw a CONFLICT error when an existing email is supplied",
		request:    UserSignUpRequest{userName: "user1", email: "user1@gmail.com", password: "test1234"},
		response:   UserSignUpResponse{},
		statusCode: 409,
		message:    "This email has already been taken",
	},
	{
		title:      "System should throw a BAD_REQUEST when a password with less than 8 characters in supplied",
		request:    UserSignUpRequest{userName: "user1", email: "user3@gmail.com", password: "test12"},
		response:   UserSignUpResponse{},
		statusCode: 400,
		message:    "Password must be at least 6 characters long",
	},
}

func TestUserSignup(t *testing.T) {

	for _, fixture := range signUpTests {
		t.Log("Executing test", fixture.title)
		server := httptest.NewServer(InitRouter())
		body, _ := json.Marshal(fixture.request)
		reader := strings.NewReader(string(body))
		request, _ := http.NewRequest("POST", "/users/new", reader)
		w := httptest.NewRecorder()
		SignUpUser(w, request)
		if w.Body.Len() == 0 {
			t.Log("Empty Response found")
			t.Fail()
		}
		server.Close()
		t.Log("---------------------------------------")
	}

}
