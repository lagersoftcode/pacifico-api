package main

import "net/http"

type Route struct {
	Name      string
	Method    string
	Pattern   string
	Authorize bool
	AdminOnly bool
	Handler   http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/CreateUser",
		true,
		true,
		CreateUser,
	},
}

func DeclareRoutes() {
	for _, route := range routes {
		handler := AddJSONResponseHeader(route.Handler)
		if route.Authorize {
			handler = Authorize.Handler(handler)
		}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}
