package routing

import (
	"net/http"
	"onehearttour/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes = Routes{
	Route{
		"Get Posts",
		"GET",
		"/posts/{id}",
		handlers.GetPosts,
	},
	Route{
		"Post Posts",
		"POST",
		"/posts",
		handlers.UploadPost,
	},
	Route{
		"Login",
		"POST",
		"/login",
		handlers.Login,
	},
	Route{
		"Signup",
		"POST",
		"/signup",
		handlers.Signup,
	},
}
