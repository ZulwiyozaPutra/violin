package controller

import (
	"net/http"

	"app/shared/view"
)

// UserSelection endpoint
func UserSelection(writer http.ResponseWriter, request *http.Request) {
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

