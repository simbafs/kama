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
}

type KamaOption func(*Kama)

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

func (k *Kama) SetPath(path string) {
	k.path = path
}

func (k *Kama) SetDevServer(devServer string) {
	u, err := url.Parse(devServer)
	if err != nil {
		panic(err)
	}
	k.devServer = u
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

func (k *Kama) Go() http.HandlerFunc {
	k.prepareFS()
	return func(w http.ResponseWriter, r *http.Request) {
		k.serveHTTP(w, r)
	}
}

func (k *Kama) Gin() gin.HandlerFunc {
	k.prepareFS()
	return func(c *gin.Context) {
		k.serveHTTP(c.Writer, c.Request)
	}
}
