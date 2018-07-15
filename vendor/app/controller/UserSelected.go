package controller

import (
	"net/http"

	"app/shared/view"
)

// UserSelected endpoint
func UserSelected(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	selectedView := view.View{
		Template: "templates/pages/select.html",
		Writer:   writer,
		Request:  request,
	}

	selectedView.Render(view.Context{
		Title:  "Your Preferred Animal",
		Answer: request.Form.Get("animalSelect"),
	})

}
