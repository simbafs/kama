package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/simbafs/kama"
)

//go:embed all:static
var static embed.FS

func main() {
	mux := http.NewServeMux()

	k := kama.New(static, kama.WithTree("/tree"))

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello, world!"}`))
	})
	mux.HandleFunc("/", k.Go())

	log.Fatal(http.ListenAndServe(":3000", mux))
}
