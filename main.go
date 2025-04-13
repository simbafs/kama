package kama

import (
	"io/fs"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Kama struct {
	static    fs.FS // the directory contain static files in production mode
	fs        fs.FS // a overlay fs
	path      string
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
	var err error
	k.fs, err = fs.Sub(k.static, k.path)
	if err != nil {
		panic(err)
	}
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
