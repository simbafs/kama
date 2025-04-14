//go:build !dev

package kama

import (
	"net/http"

	"github.com/simbafs/kama/httpfs"
)

// func tree(fs afero.Fs, w io.Writer) {
// 	afero.Walk(fs, ".", func(path string, info os.FileInfo, err error) error {
// 		depth := strings.Count(strings.TrimPrefix(path, "/"), "/")
// 		indentation := strings.Repeat("│   ", depth)
//
// 		if info.IsDir() {
// 			fmt.Fprintf(w, "%s├── %s/\n", indentation, info.Name())
// 		} else {
// 			fmt.Fprintf(w, "%s├── %s\n", indentation, info.Name())
// 		}
//
// 		return nil
// 	})
// }

func (k *Kama) serveHTTP(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path == "/tree" {
	// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// 	w.WriteHeader(http.StatusOK)
	// 	tree(k.fs, w)
	// 	return
	// }

	http.FileServer(httpfs.NewHTTPFs(k.fs)).ServeHTTP(w, r)
}
