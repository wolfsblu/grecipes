package routes

import (
	"github.com/wolfsblu/grecipes/app"
	"io/fs"
	"net/http"
)

func Assets(w http.ResponseWriter, r *http.Request) {
	sub, _ := fs.Sub(app.DistFS, "dist")
	h := http.FileServerFS(sub)
	h.ServeHTTP(w, r)
}
