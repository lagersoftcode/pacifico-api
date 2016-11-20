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
}

func DeclareRoutes() {
	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			corsHandler(w, r)
		})

	for _, route := range routes {
		handler := AddJSONResponseHeader(route.Handler)
		if route.Authorize {
			handler = Authorize(handler, route.AdminOnly)
		}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
}

func corsHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", config.CORSDomain)
	rw.Header().Set("Access-Control-Allow-Headers", "Content-type, With-Credentials")
	rw.WriteHeader(http.StatusOK)
}
