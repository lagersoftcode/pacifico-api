package main

import (
	"encoding/base64"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

func AddJSONResponseHeader(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		inner.ServeHTTP(rw, req)
	})
}

var Authorize = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		decoded, err := base64.URLEncoding.DecodeString(config.TokenKey)
		if err != nil {
			return nil, err
		}
		return decoded, nil
	},
})
