package routing

import (
	"github.com/wolfsblu/go-chef/webapp"
	"io/fs"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	sub, _ := fs.Sub(app.DistFS, "dist")
	http.ServeFileFS(w, r, sub, "index.html")
}

func assets(w http.ResponseWriter, r *http.Request) {
	sub, _ := fs.Sub(app.DistFS, "dist")
	h := http.FileServerFS(sub)
	h.ServeHTTP(w, r)
}
