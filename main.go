// Package kama handle serving static file in proudction mode and development mode
// In production mode, it serves the static files from the embedded and overlay filesystem
// In development mode, it proxies the requests to the development server(e.g. Vite, Next.js dev server)
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
	static      fs.FS    // the directory contain static files in production mode
	staticPath  string   // path to static directory, see [WithPath] below
	overlayPath string   // path to overlay directory in current working directory
	fs          afero.Fs // a overlay fs
	devServer   *url.URL
	tree        string
}

type KamaOption func(*Kama) error

// New creates a new Kama instance
func New(f fs.FS, devServer string, opts ...KamaOption) (*Kama, error) {
	dev, err := url.Parse(devServer)
	if err != nil {
		return nil, err
	}

	k := &Kama{
		static:     f,
		staticPath: "static",
		devServer:  dev,
	}
	for _, opt := range opts {
		if err := opt(k); err != nil {
			return nil, err
		}
	}
	return k, nil
}

// WithStaticPath sets the path to the static directory relative to the current working directory
// This option is used to remove the directory nesting if needed.
// For example, if you load the embed fs like this
//
//	//go:embed all:fromtend
//	//var staticFS embed.FS
//
// The strucuture of the staticFS look like this:
//
//	static
//	 ├── index.html
//	 └── style.css
//
// With this option, it will be possible to serve the files correctly by setting the path to "frontend"
//
//	/ -> frontend/index.html
//	/style.css -> fromtend/style.css
//
// If not set, it will be
//
//	/static -> frontend/index.html
//	/static/style.css -> fromtend/style.css
func WithStaticPath(path string) KamaOption {
	return func(k *Kama) error {
		k.staticPath = path
		return nil
	}
}

// WithOverlayPath sets the path to the overlay directory in the current working directory
// This directory is prioritized over the embedded static files.
func WithOverlayPath(path string) KamaOption {
	return func(k *Kama) error {
		k.overlayPath = path
		return nil
	}
}

// WithTree enable a endpoint to show the tree of filesystem
// This is useful to inspect if the [WithOverlayPath] is working correctly
func WithTree(tree string) KamaOption {
	return func(k *Kama) error {
		k.tree = tree
		return nil
	}
}

func (k *Kama) prepareFS() {
	embed, err := fs.Sub(k.static, k.staticPath)
	if err != nil {
		panic(err)
	}

	if k.overlayPath == "" {
		k.fs = afero.FromIOFS{FS: embed}
		return
	}

	overlay := os.DirFS(k.overlayPath)
	k.fs = overlayfs.New(overlayfs.Options{
		Fss: []afero.Fs{afero.FromIOFS{FS: overlay}, afero.FromIOFS{FS: embed}},
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
