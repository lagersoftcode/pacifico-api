package main

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const authCookie = "auth_token"

func ValidateToken(res http.ResponseWriter, req *http.Request) *jwt.Token {
	defToken := jwt.Token{Valid: false}
	cookie, err := req.Cookie(authCookie)

	if err != nil {
		return &defToken
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte(config.AuthKey), nil
	})

	if err != nil {
		return &defToken
	}

	return token
}

func SetToken(rw http.ResponseWriter, userName string, isAdmin bool) {
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	expireCookie := time.Now().Add(time.Hour * 24)

	claims := Claims{
		userName,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "pacific-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(config.AuthKey))
	cookie := http.Cookie{Name: authCookie, Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(rw, &cookie)
}
