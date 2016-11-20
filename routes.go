package main

import "net/http"

type Route struct {
	Name      string
	Method    string
	Pattern   string
	Authorize bool
	AdminOnly bool
	Handler   BaseRouteHandler
}

type Routes []Route

var routes = Routes{
	Route{
		"Login",
		"POST",
		"/Login",
		false,
		false,
		Login,
	},
	Route{
		"CreateUser",
		"POST",
		"/CreateUser",
		true,
		true,
		CreateUser,
	},
	Route{
		"GetUsers",
		"Get",
		"/GetUsers",
		true,
		false,
		GetUsers,
	},
}

func DeclareRoutes() {
	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			CorsHandler(w, r)
		})

	for _, route := range routes {
		handler := BaseHandler(route.Handler, route.Authorize, route.AdminOnly)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
}
