package main

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string
	UserId   string
	IsAdmin  bool
	jwt.StandardClaims
}

type BaseRouteHandler func(rw http.ResponseWriter, r *http.Request, routeData RouteData)
type RouteData struct {
	Body     []byte
	Username string
	UserId   string
}

type Response struct {
	HttpResult int
}

// Login

type CreateUserRequest struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Email     string
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

type GetTrophiesResponse struct {
	Trophies []Trophy
}

// Medals

type CreateMedalRequest struct {
	Name        string
	Image       string
	Material    string
	Description string
	ScoreAmount uint
}

type GetMedalsResponse struct {
	Medals []Medal
}

// Actions

type GiveTrophyRequest struct {
	UserId   string
	TrophyId string
}

type GiveMedalRequest struct {
	UserId  string
	MedalId string
}

type PublicAction struct {
	CreatedAt  time.Time
	SourceUser string
	TargetUser string
	Item       string
}

type GetLastActionsResponse struct {
	LastActions []PublicAction
}
