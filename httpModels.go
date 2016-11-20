package main

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string
	IsAdmin  bool
	jwt.StandardClaims
}

type AuthorizedRequest struct {
	AuthToken string
}

type BaseRouteHandler func(rw http.ResponseWriter, r *http.Request, routeData RouteData)
type RouteData struct {
	Body []byte
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
	Success   bool
	Message   string
	AuthToken string
}
