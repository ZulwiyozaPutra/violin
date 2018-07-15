package controller

import (
	"fmt"
	"net/http"
)

// Index endpoint
func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the violin")
}
