package api

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
