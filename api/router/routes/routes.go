package routes

import (
	"go_jwt_auth/api/middlewares"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type Routables interface {
	Routes() []Route
}



type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	Auth    Auth
}

type Auth struct {
	UserRequired  bool
	AdminRequired bool
}

func Load(routables []Routables) []Route {
	allRoutes := []Route{}
	for _, router := range routables {
		allRoutes = append(allRoutes, router.Routes()...)
	}
	return allRoutes
}

func SetupRoutes(r *mux.Router, routables []Routables, authconfig middlewares.AuthConfig) *mux.Router {
	// routes := Load(routables)
	// for _, route := range routes {
	// 	r.HandleFunc(route.Uri,
	// 		middlewares.SetMiddlewareLogger(
	// 			middlewares.SetMiddlewareJSON(route.Handler)),
	// 	).Methods(route.Method)
	// }

	routes := Load(routables)
	setMiddlewares := []middlewares.Middleware{
		middlewares.SetMiddlewareLogger,
		middlewares.SetMiddlewareJSON,
	}

	setAdminMiddlewares := []middlewares.Middleware{
		middlewares.SetMiddlewareLogger,
		middlewares.SetMiddlewareJSON,
		middlewares.SetMiddlewareAuth(authconfig),

	}
	for _, route := range routes {

		if route.Auth.AdminRequired {
			r.HandleFunc(route.Uri,
				middlewares.BuildMiddlewareChain(
					route.Handler,
					setAdminMiddlewares,
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri,
				middlewares.BuildMiddlewareChain(
					route.Handler,
					setMiddlewares,
				),
			).Methods(route.Method)
		}

	
	}
	return r
}


func reverseAny(s interface{}) {
    n := reflect.ValueOf(s).Len()
    swap := reflect.Swapper(s)
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        swap(i, j)
    }
}