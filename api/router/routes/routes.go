package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matkinhig/go-blogs/api/middlewares"
)

type Route struct {
	Uri          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Load() []Route {
	routes := userRoutes
	routes = append(routes, postRoutes...)
	routes = append(routes, loginRoutes...)
	routes = append(routes, categoryRoutes...)
	routes = append(routes, productRoutes...)
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		if route.AuthRequired {
			r.HandleFunc(route.Uri, middlewares.SetMiddlewareLogger(
				middlewares.SetMiddlewareJSON(
					middlewares.SetMiddlewareAuthentication(route.Handler))),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.SetMiddlewareLogger(
				middlewares.SetMiddlewareJSON(route.Handler)),
			).Methods(route.Method)
		}
	}
	return r
}
