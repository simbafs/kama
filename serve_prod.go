//go:build !dev

package kama

import (
	"net/http"
)

func (k *Kama) serveHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(k.fs)).ServeHTTP(w, r)
}
