package routes

import (
	"github.com/wolfsblu/grecipes/app"
	"io/fs"
	"net/http"
)

func app(w http.ResponseWriter, r *http.Request) {
	sub, _ := fs.Sub(app.DistFS, "dist")
	http.ServeFileFS(w, r, sub, "index.html")
}
