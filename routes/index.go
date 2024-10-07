package routes

import (
	"context"
	"github.com/wolfsblu/grecipes/templates"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	c := templates.Index()
	_ = c.Render(context.Background(), w)
}
