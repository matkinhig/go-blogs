package router

import (
	"github.com/gorilla/mux"
	"github.com/matkinhig/go-blogs/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(r)
}
