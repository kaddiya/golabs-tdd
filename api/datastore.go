package api

import "errors"

var userDataStore = []user{}

//UserDao exposes the methods to be able to store everything in the database
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
