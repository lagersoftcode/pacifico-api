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
	Route{"Login", "POST", "/Login", false, false, Login},
	Route{"CreateUser", "POST", "/CreateUser", true, true, CreateUser},
	Route{"GetUsers", "GET", "/GetUsers", true, false, GetUsers},
	Route{"CreateTrophy", "POST", "/CreateTrophy", true, true, CreateTrophy},
	Route{"GetTrophies", "GET", "/GetTrophies", true, false, GetTrophies},
	Route{"CreateMedal", "POST", "/CreateMedal", true, true, CreateMedal},
	Route{"GetMedals", "GET", "/GetMedals", true, false, GetMedals},
	Route{"GiveTrophy", "POST", "/GiveTrophy", true, true, GiveTrophy},
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
