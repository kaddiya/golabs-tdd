package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

//UserSignUpRequest encapuslates the signup request
type UserSignUpRequest struct {
	userName string
	email    string
	password string
}

//UserSignUpResponse encapsulates the Signup response
type UserSignUpResponse struct {
	userName string
	userID   int
}

//SignUpUser exposes the signup API
func SignUpUser(w http.ResponseWriter, r *http.Request) {

}

//InitRouter initialises the mux router
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/new", SignUpUser).Methods("POST")
	return r
}
