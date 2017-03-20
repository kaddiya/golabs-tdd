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

const userIDValue = 0

//SignUpUser exposes the signup API
func SignUpUser(w http.ResponseWriter, r *http.Request) {
	reqBody := &UserSignUpRequest{}
	parseRequestBody(r.Body, reqBody)

	//check for valid email address
	if !strings.Contains(reqBody.Email, "@") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//check for the length of the password
	if len(reqBody.Password) < 8 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

//InitRouter initialises the mux router
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/new", SignUpUser).Methods("POST")
	return r
}

func parseRequestBody(r io.Reader, target interface{}) interface{} {
	return json.NewDecoder(r).Decode(target)
}

//functions to populate the DataStore
