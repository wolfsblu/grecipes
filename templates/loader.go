package templates

import (
	"html/template"
)

var Template map[string]*template.Template

func init() {
	if Template == nil {
		Template = make(map[string]*template.Template)
	}
	Template["index.html"] = template.Must(
		template.ParseFiles("templates/base.html", "templates/index.html"),
	)
}
