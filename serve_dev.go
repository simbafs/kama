//go:build dev

package kama

import (
	"net/http"
	"net/http/httputil"
)

func (k *Kama) serveHTTP(w http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(k.devServer)
	proxy.ServeHTTP(w, r)
}
