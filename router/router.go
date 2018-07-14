package router

import (
	"net/http"
)

// Route object
type Route struct {
	Path    string
	Handler func(writer http.ResponseWriter, request *http.Request)
}

// HandleRoutes with parameter a slice of routes
func HandleRoutes(routes []Route) {
	for i := 0; i < len(routes); i++ {
		route := routes[i]
		http.HandleFunc(route.Path, route.Handler)
	}
}
