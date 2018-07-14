package main

import (
	"fmt"
	"net/http"

	"github.com/zulwiyozaputra/violin/view"
)

// index
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the violin")
}

// UserSelection endpoint
func userSelection(writer http.ResponseWriter, request *http.Request) {
	selectionView := view.View{
		Template: "templates/pages/select.html",
		Writer:   writer,
		Request:  request,
	}

	selectionView.Render(view.Context{
		Title: "Which Do You Prefer?",
		RadioButtons: []view.RadioButton{
			view.RadioButton{
				Name:       "animalSelect",
				Value:      "CATS",
				IsDisabled: false,
				IsChecked:  false,
				Text:       "Cats",
			},
			view.RadioButton{
				Name:       "animalSelect",
				Value:      "DOGS",
				IsDisabled: false,
				IsChecked:  false,
				Text:       "Dogs",
			},
		},
	})
}

// UserSelected endpoint
func userSelected(writer http.ResponseWriter, request *http.Request) {
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
