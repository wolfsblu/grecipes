package routes

import (
	"github.com/wolfsblu/grecipes/templates"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	_ = templates.Template["index.html"].ExecuteTemplate(w, "base.html", nil)
}
