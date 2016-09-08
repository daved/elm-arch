package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/codemodus/mixmux"
)

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("assets", r.URL.Path[7:]))
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("assets/icon/"+r.URL.Path))
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("html", r.URL.Path))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "" {
		http.ServeFile(w, r, filepath.Join("html", "index.html"))
		return
	}

	if len(r.URL.Path) < 5 {
		http.ServeFile(w, r, r.URL.Path+".html")
		return
	}

	sfx := r.URL.Path[len(r.URL.Path)-5:]

	if sfx == ".html" || sfx[1:] == ".htm" {
		htmlHandler(w, r)
		return
	}

	if sfx[1:] == ".ico" {
		iconHandler(w, r)
		return
	}

	http.ServeFile(w, r, r.URL.Path+".html")
}

func main() {
	mOpts := &mixmux.Options{NotFound: http.HandlerFunc(indexHandler)}
	m := mixmux.NewRouter(mOpts)

	m.Get("/assets/*x", http.HandlerFunc(assetsHandler))

	if err := http.ListenAndServe(":4001", m); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
