//go:build !dev

package kama

import (
	"io/fs"
	"log"
	"net/http"
)

func (k *Kama) serveHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := fs.Sub(k.fs, k.path)
	if err != nil {
		log.Printf("Oops, there's an error: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.FileServer(http.FS(f)).ServeHTTP(w, r)
}
