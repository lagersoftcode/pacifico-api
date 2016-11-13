package main

type Response struct {
	HttpResult int
}

type CreateUserRequest struct {
	Username string
	Password string
	Token    string
}

type CreateUserResponse struct {
	Response
}
