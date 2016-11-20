package main

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func ValidateToken(res http.ResponseWriter, req AuthorizedRequest) *jwt.Token {
	defToken := jwt.Token{Valid: false}

	token, err := jwt.ParseWithClaims(string(req.AuthToken), &Claims{}, func(token *jwt.Token) (interface{}, error) {
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

func GetToken(rw http.ResponseWriter, userName string, isAdmin bool) string {
	expireToken := time.Now().Add(time.Hour * 24).Unix()

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
	return signedToken
}
