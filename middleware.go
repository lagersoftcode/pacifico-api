package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BaseHandler(handler BaseRouteHandler, requiresAuth bool, adminOnly bool) http.Handler {

	var routeData RouteData

	httpHandler := addJSONResponseHeader(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, routeData)
	}))

	return Authorize(httpHandler, &routeData, requiresAuth, adminOnly)
}

func addJSONResponseHeader(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", config.CORSDomain)
		rw.Header().Set("Access-Control-Allow-Headers", "Content-type, With-Credentials")
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		inner.ServeHTTP(rw, req)
	})
}

func Authorize(handler http.Handler, routeData *RouteData, requiresAuth bool, adminOnly bool) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		body, readErr := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if readErr != nil {
			panic(readErr)
		}

		routeData.Body = body

		if !requiresAuth {
			handler.ServeHTTP(res, req)
			return
		}

		var request AuthorizedRequest
		parseErr := json.Unmarshal(routeData.Body, &request)

		if parseErr != nil {
			panic(parseErr)
		}

		token := ValidateToken(res, request)

		claims, claimsOk := token.Claims.(*Claims)
		if token.Valid && claimsOk && (!adminOnly || adminOnly && claims.IsAdmin) {
			ctx := context.WithValue(req.Context(), "claims", *claims)
			handler.ServeHTTP(res, req.WithContext(ctx))
			return
		}

		http.NotFound(res, req)
	})
}

func CorsHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", config.CORSDomain)
	rw.Header().Set("Access-Control-Allow-Headers", "Content-type, With-Credentials")
	rw.WriteHeader(http.StatusOK)
}
