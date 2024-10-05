package templates

import (
	"embed"
	"html/template"
)

//go:embed assets
var AssetsFS embed.FS

//go:embed *.html
var TemplateFS embed.FS

var Template map[string]*template.Template

func init() {
	if Template == nil {
		Template = make(map[string]*template.Template)
	}
	Template["index.html"] = template.Must(
		template.ParseFS(TemplateFS, "base.html", "index.html"),
	)
}
