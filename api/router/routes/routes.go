package routes

import (
	"go_jwt_auth/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Routables interface {
	Routes() []Route
}

type Route struct {
	Uri          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Load(routables []Routables) []Route {
	allRoutes := []Route{}
	for _, router := range routables {
		allRoutes = append(allRoutes, router.Routes()...)
	}
	return allRoutes
}

func SetupRoutes(r *mux.Router, routables []Routables) *mux.Router {
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
	for _, route := range routes {
		r.HandleFunc(route.Uri,
			middlewares.BuildMiddlewareChain(
				route.Handler,
				setMiddlewares,
			),
		).Methods(route.Method)
	}
	return r
}
