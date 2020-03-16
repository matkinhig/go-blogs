package routes

import (
	"net/http"

	"github.com/matkinhig/go-blogs/api/controllers"
)

var productRoutes = []Route{
	Route{
		Uri:          "/products",
		Method:       http.MethodPost,
		Handler:      controllers.CreateProduct,
		AuthRequired: false,
	},
}
