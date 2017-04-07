package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//UserSignUpRequest encapuslates the signup request
type UserSignUpRequest struct {
	UserName string
	Email    string
	Password string
}

//UserSignUpResponse encapsulates the Signup response
type UserSignUpResponse struct {
	UserName string
	UserID   int
}

type user struct {
	UserID   int
	UserName string
	Email    string
	Password string
}

//SignUpUser exposes the signup API
func (c *Container) SignUpUser(w http.ResponseWriter, r *http.Request) {
	reqBody := &UserSignUpRequest{}
	parseRequestBody(r.Body, reqBody)

	//check for valid email address
	if !strings.Contains(reqBody.Email, "@") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please enter a valid email id"))
		return
	}
	//check for the length of the password
	if len(reqBody.Password) < 8 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password must be at least 8 characters long"))
		return
	}

	_, uniqueViolationErr := c.Dao.isEmailIDUnique(reqBody.Email)

	if uniqueViolationErr != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("This email has already been taken"))
		return
	}

	c.Dao.saveUser(reqBody)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You have successfully signed up"))
}

//InitRouter initialises the mux router
func (c *Container) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/new", c.SignUpUser).Methods("POST")
	return r
}

func parseRequestBody(r io.Reader, target interface{}) interface{} {
	return json.NewDecoder(r).Decode(target)
}
