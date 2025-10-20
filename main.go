// This package provides HTTP server handlers for serving static files.
// In development mode, it proxies requests to a frontend development server.
// In production mode, it serves files from the embedded filesystem,
// with optional overrides from a local directory in the current working directory.
package kama

import (
	"io/fs"
	"net/http"
	"net/url"
	"os"

	"github.com/bep/overlayfs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

type Kama struct {
	static    fs.FS    // the directory contain static files in production mode
	fs        afero.Fs // a overlay fs
	path      string   // path to [static] directory
	devServer *url.URL
	tree      string
}

type KamaOption func(*Kama)

// New creates a new Kama instance
func New(f fs.FS, opts ...KamaOption) *Kama {
	k := &Kama{
		static: f,
		path:   "static",
		devServer: &url.URL{
			Scheme: "http",
			Host:   "localhost:3001",
		},
	}
	for _, opt := range opts {
		opt(k)
	}
	return k
}

// WithPath sets the path to the static directory (default is "static")
func WithPath(path string) KamaOption {
	return func(k *Kama) {
		k.path = path
	}
}

// WithDevServer sets the dev server URL (default is http://localhost:3001)
func WithDevServer(devServer string) KamaOption {
	return func(k *Kama) {
		u, err := url.Parse(devServer)
		if err != nil {
			panic(err)
		}
		k.devServer = u
	}
}

// WithTree enable a endpoint to show the tree of filesystem
func WithTree(tree string) KamaOption {
	return func(k *Kama) {
		k.tree = tree
	}
}

func (k *Kama) prepareFS() {
	A, err := fs.Sub(k.static, k.path)
	if err != nil {
		panic(err)
	}

	B := os.DirFS(k.path)
	k.fs = overlayfs.New(overlayfs.Options{
		Fss: []afero.Fs{afero.FromIOFS{FS: B}, afero.FromIOFS{FS: A}},
	})
}

// Go returns a http.HandlerFunc for http mux in standard library
func (k *Kama) Go() http.HandlerFunc {
	k.prepareFS()
	return func(w http.ResponseWriter, r *http.Request) {
		k.serveHTTP(w, r)
	}
}

// Gin returns a gin.HandlerFunc for gin
func (k *Kama) Gin() gin.HandlerFunc {
	k.prepareFS()
	return func(c *gin.Context) {
		k.serveHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
