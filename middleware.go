package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AddJSONResponseHeader(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		inner.ServeHTTP(rw, req)
	})
}

func Authorize(handler http.Handler, adminOnly bool) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		cookie, err := req.Cookie("auth_token")
		if err != nil {
			http.NotFound(res, req)
			return
		}
		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte(config.AuthKey), nil
		})

		if err != nil {
			http.NotFound(res, req)
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			if adminOnly && claims.IsAdmin {
				ctx := context.WithValue(req.Context(), "claims", *claims)
				handler.ServeHTTP(res, req.WithContext(ctx))
				return
			}

			http.NotFound(res, req)
			return
		}
	})
}
