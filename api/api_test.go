package api

var signUpTests = []struct {
	request    UserSignUpRequest
	response   UserSignUpResponse
	statusCode int
	message    string
}{{
	request:    UserSignUpRequest{userName: "user1", email: "user1@gmail.com", password: "test1234"},
	response:   UserSignUpResponse{userName: "user1", userID: 1},
	statusCode: 200,
	message:    "You have successfully signed up",
},
	{
		request:    UserSignUpRequest{userName: "user1", email: "user1", password: "test1234"},
		response:   UserSignUpResponse{},
		statusCode: 400,
		message:    "Please enter a valid email id",
	},
	{
		request:    UserSignUpRequest{userName: "user1", email: "user1@gmail.com", password: "test1234"},
		response:   UserSignUpResponse{},
		statusCode: 409,
		message:    "This email has already been taken",
	},
	{
		request:    UserSignUpRequest{userName: "user1", email: "user1", password: "test12"},
		response:   UserSignUpResponse{},
		statusCode: 409,
		message:    "Password must be at least 6 characters long",
	},
}
