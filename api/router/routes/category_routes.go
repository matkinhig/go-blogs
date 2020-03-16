package routes

import (
	"net/http"

	"github.com/matkinhig/go-blogs/api/controllers"
)

var categoryRoutes = []Route{
	Route{
		Uri:          "/category",
		Method:       http.MethodPost,
		Handler:      controllers.CreateCategory,
		AuthRequired: false,
	},
	// Route{
	// 	Uri: "/category",
	// 	Method: http.MethodPost,
	// 	Handler: controllers.CreateCategory,
	// 	AuthRequired: false,
	// },
	// Route{
	// 	Uri: "/category",
	// 	Method: http.MethodPost,
	// 	Handler: controllers.CreateCategory,
	// 	AuthRequired: false,
	// },
	// Route{
	// 	Uri: "/category",
	// 	Method: http.MethodPost,
	// 	Handler: controllers.CreateCategory,
	// 	AuthRequired: false,
	// },
}
