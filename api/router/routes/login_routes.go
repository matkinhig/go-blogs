package routes

import (
	"net/http"

	"github.com/matkinhig/go-blogs/api/controllers"
)

var loginRoutes = []Route{
	Route{
		Uri:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
