package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateUser(rw http.ResponseWriter, req *http.Request) {

	body, readErr := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if readErr != nil {
		panic(readErr)
	}

	var request CreateUserRequest
	parseErr := json.Unmarshal(body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	response := CreateUserResponse{Response{http.StatusCreated}}

	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

func Login(rw http.ResponseWriter, req *http.Request) {

	body, readErr := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if readErr != nil {
		panic(readErr)
	}

	var request LoginRequest
	parseErr := json.Unmarshal(body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	var response LoginResponse
	var user User
	db.Where(&User{UserName: request.Username}).First(&user)

	if user.ID == 0 {
		response.Success = false
		response.Message = "Invalid credentials"
		rw.WriteHeader(http.StatusUnauthorized)
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			response.Success = false
			response.Message = "Invalid credentials"
			rw.WriteHeader(http.StatusUnauthorized)
		} else {
			response.Success = true
			setToken(rw, user.UserName, user.IsAdmin)
			rw.WriteHeader(http.StatusOK)
		}
	}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

func setToken(rw http.ResponseWriter, userName string, isAdmin bool) {
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	expireCookie := time.Now().Add(time.Hour * 24)

	claims := Claims{
		userName,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "x",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(config.TokenKey))
	cookie := http.Cookie{Name: "auth_token", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(rw, &cookie)
}
