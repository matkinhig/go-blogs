package routes

import (
	"net/http"

	controller "github.com/matkinhig/go-blogs/api/controllers"
)

var userRoutes = []Route{
	Route{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controller.GetUsers,
	},
	Route{
		Uri:     "/users",
		Method:  http.MethodPost,
		Handler: controller.CreateUser,
	},
	Route{
		Uri:     "/users/{id}",
		Method:  http.MethodGet,
		Handler: controller.GetUser,
	},
	Route{
		Uri:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controller.UpdateUser,
	},
	Route{
		Uri:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: controller.DeleteUser,
	},
}
