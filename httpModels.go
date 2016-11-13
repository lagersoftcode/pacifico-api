package main

import jwt "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string
	IsAdmin  bool
	jwt.StandardClaims
}

type Response struct {
	HttpResult int
}

type CreateUserRequest struct {
	Username string
	Password string
}

type CreateUserResponse struct {
	Response
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Success bool
	Message string
}
