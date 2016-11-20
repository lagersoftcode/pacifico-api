package main

import (
	"context"
	"net/http"
)

func AddJSONResponseHeader(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", config.CORSDomain)
		rw.Header().Set("Access-Control-Allow-Headers", "Content-type, With-Credentials")
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		inner.ServeHTTP(rw, req)
	})
}

func Authorize(handler http.Handler, adminOnly bool) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		token := ValidateToken(res, req)

		claims, claimsOk := token.Claims.(*Claims)
		if token.Valid && claimsOk && (!adminOnly || adminOnly && claims.IsAdmin) {
			ctx := context.WithValue(req.Context(), "claims", *claims)
			handler.ServeHTTP(res, req.WithContext(ctx))
			return
		}

		http.NotFound(res, req)
	})
}
