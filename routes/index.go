package routes

import (
	"context"
	"github.com/wolfsblu/grecipes/templates/pages"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	c := pages.Index()
	_ = c.Render(context.Background(), w)
}
