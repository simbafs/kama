package kama

import (
	"io/fs"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Kama struct {
	fs        fs.FS
	path      string
	devServer *url.URL
}

type KamaOption func(*Kama)

func New(f fs.FS, opts ...KamaOption) *Kama {
	k := &Kama{
		fs:   f,
		path: "static",
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

func (k *Kama) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	k.serveHTTP(w, r)
}

func (k *Kama) Gin() gin.HandlerFunc {
	return func(c *gin.Context) {
		k.serveHTTP(c.Writer, c.Request)
	}
}
