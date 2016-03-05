package server

import (
	"net/http"

	. "github.com/showntop/circle-core/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	//嵌套路由的问题
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Home{}.Index},

	Route{"UsersIndex", "GET", "/users", User{}.Index},
	Route{"UsersShow", "GET", "/users/:id", User{}.Show},
	Route{"UsersCreate", "POST", "/users", User{}.Create},
}
