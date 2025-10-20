//go:build dev

package kama

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func (k *Kama) serveHTTP(w http.ResponseWriter, r *http.Request) {
	if k.tree != "" && r.URL.Path == k.tree {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "Tree is only available in the production mode.")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(k.devServer)
	proxy.ServeHTTP(w, r)
}
