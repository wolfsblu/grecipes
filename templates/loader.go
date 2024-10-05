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
		template.ParseFS(TemplateFS, "html/base.html", "html/index.html"),
	)
}
