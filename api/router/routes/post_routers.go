package routes

import (
	"net/http"

	"github.com/matkinhig/go-blogs/api/controllers"
)

var postRoutes = []Route{
	Route{
		Uri:          "/posts",
		Method:       http.MethodPost,
		Handler:      controllers.CreatePost,
		AuthRequired: true,
	},
	Route{
		Uri:          "/posts",
		Method:       http.MethodGet,
		Handler:      controllers.GetPosts,
		AuthRequired: true,
	},
	Route{
		Uri:          "/posts/{id}",
		Method:       http.MethodGet,
		Handler:      controllers.GetPost,
		AuthRequired: false,
	},
	Route{
		Uri:          "/posts/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdatePost,
		AuthRequired: true,
	},
	Route{
		Uri:          "/posts/{id}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeletePost,
		AuthRequired: true,
	},
}
