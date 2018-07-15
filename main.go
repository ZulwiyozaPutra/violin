package main

import (
	"net/http"
	"os"

	"app/shared/router"
	"app/route"
)



func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8080"
}

func main() {
	port := getPort()
	router.HandleRoutes(route.Routes)
	http.ListenAndServe(port, nil)
}
