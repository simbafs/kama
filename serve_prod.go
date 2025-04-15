//go:build !dev

package kama

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/simbafs/kama/httpfs"
	"github.com/spf13/afero"
)

func tree(fs afero.Fs, w io.Writer) {
	afero.Walk(fs, ".", func(path string, info os.FileInfo, err error) error {
		depth := strings.Count(strings.TrimPrefix(path, "/"), "/")
		indentation := strings.Repeat("│   ", depth)

		if info.IsDir() {
			fmt.Fprintf(w, "%s├── %s/\n", indentation, info.Name())
		} else {
			fmt.Fprintf(w, "%s├── %s\n", indentation, info.Name())
		}

		return nil
	})
}

func (k *Kama) serveHTTP(w http.ResponseWriter, r *http.Request) {
	if k.tree != "" && r.URL.Path == k.tree {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		tree(k.fs, w)
		return
	}

	http.FileServer(httpfs.NewHTTPFs(k.fs)).ServeHTTP(w, r)
}
