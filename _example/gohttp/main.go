package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/simbafs/kama/v2"
)

//go:embed all:static
var static embed.FS

func main() {
	mux := http.NewServeMux()

	k, _ := kama.New(static, "http://localhost:3001", kama.WithTree("/tree"), kama.WithOverlayPath("overlay"))

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello, world!"}`))
	})
	mux.HandleFunc("/", k.Go())

	log.Fatal(http.ListenAndServe(":3000", mux))
}
