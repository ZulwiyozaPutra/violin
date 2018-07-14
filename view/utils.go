package view

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// View object
type View struct {
	Template string
	Writer   http.ResponseWriter
	Request  *http.Request
}

// Context for any rendered html
type Context struct {
	Title        string
	RadioButtons []RadioButton
	Answer       string
}

// Render the view
func (view View) Render(context Context) {
	template, err := template.ParseFiles(view.Template)
	logError("Template parsing error: ", err)
	err = template.Execute(view.Writer, context)
	logError("Template executing error: ", err)
}

func logError(message string, err error) {
	if err != nil {
		log.Print(message, err)
		os.Exit(3)
	}
}
