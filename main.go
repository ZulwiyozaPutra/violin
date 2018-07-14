package main

import (
	"net/http"
	"os"

	"github.com/zulwiyozaputra/violin/router"
)

// ROUTES consists of slices of Routes
var ROUTES = []router.Route{
	router.Route{Path: "/", Handler: index},
	router.Route{Path: "/select", Handler: userSelection},
	router.Route{Path: "/selected", Handler: userSelected},
}

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8080"
}

func main() {
	port := getPort()
	router.HandleRoutes(ROUTES)
	http.ListenAndServe(port, nil)
}
