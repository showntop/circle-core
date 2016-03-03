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
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Home{}.Index},
}
