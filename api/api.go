package api

import (
	"encoding/json"
	"errors"
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

var userIDValue = 0

var userDataStore = []user{}

var dao UserDao

func SetUserDao(passedDao UserDao) {
	dao = passedDao
}

//SignUpUser exposes the signup API
func SignUpUser(w http.ResponseWriter, r *http.Request) {
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
	//var dao UserDao = &InMemoryUserDao{}
	_, uniqueViolationErr := dao.isEmailIDUnique(reqBody.Email)

	if uniqueViolationErr != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("This email has already been taken"))
		return
	}

	dao.saveUser(reqBody)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You have successfully signed up"))
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

//UserDao will be the interface for all user related functions
type UserDao interface {
	isEmailIDUnique(email string) (bool, error)
	saveUser(u *UserSignUpRequest)
}

//InMemoryUserDao handles the user populationg mechanism in memory
type InMemoryUserDao struct {
}

//make InMemoryUserDao implement userDao
func (dao *InMemoryUserDao) isEmailIDUnique(email string) (bool, error) {

	for _, user := range userDataStore {
		if user.Email == email {
			return false, errors.New("This Email Id has already been taken")
		}
	}
	return true, nil
}

func (dao *InMemoryUserDao) saveUser(u *UserSignUpRequest) {
	userIDValue = userIDValue + 1
	newUser := user{UserID: userIDValue, UserName: u.UserName, Email: u.Email, Password: u.Password}
	userDataStore = append(userDataStore, newUser)
}
