package main

import (
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
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		inner.ServeHTTP(rw, req)
	})
}

func Authorize(handler http.Handler, routeData *RouteData, requiresAuth bool, adminOnly bool) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		res.Header().Set("Access-Control-Allow-Origin", config.CORSDomain)
		res.Header().Set("Access-Control-Allow-Headers", "Content-type, With-Credentials")

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

		var tokenInQuery = req.URL.Query().Get("AuthToken")

		token := ValidateToken(tokenInQuery)

		claims, claimsOk := token.Claims.(*Claims)
		if token.Valid && claimsOk && (!adminOnly || adminOnly && claims.IsAdmin) {
			routeData.Username = claims.Username
			handler.ServeHTTP(res, req)
			return
		}

		res.WriteHeader(http.StatusUnauthorized)
		return
	})
}

func CorsHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", config.CORSDomain)
	rw.Header().Set("Access-Control-Allow-Headers", "Content-type, With-Credentials")
	rw.WriteHeader(http.StatusOK)
}
