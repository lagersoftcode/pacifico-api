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

type BaseRouteHandler func(rw http.ResponseWriter, r *http.Request, routeData RouteData)
type RouteData struct {
	Body []byte
}

type Response struct {
	HttpResult int
}

// Login

type CreateUserRequest struct {
	Username string
	Password string
}

type CreateUserResponse struct {
	Response
}

type GetUsersResponse struct {
	Users []PublicUser
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

// Trophies

type CreateTrophyRequest struct {
	Name        string
	Image       string
	Description string
	ScoreAmount uint
}

// Medals

type CreateMedalRequest struct {
	Name        string
	Image       string
	Material    int
	Description string
	ScoreAmount uint
}
