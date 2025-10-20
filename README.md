# kama

> /kam-á/  
> 橘子的臺語  
> Tangerine in Taiwanese Hokkien

Kama helps you deal with static in http server. In development mode, it redirect the request to dev server(such as Nextjs, Astro etc). In production mode, it use go embed package to serve compiled static file.

## Features

- Supports both `net/http` and `gin` framework
- Proxies requests to a frontend dev server during development
- Serves static files via `embed.FS` in production
- Allows local static file override on top of embedded files

## Installation

```bash
go get github.com/simbafs/kama
```

# Quick Overview

```go
//go:embed all:static
var embededFS embed.FS

k := kama.New(embeddedFS,
  kama.WithDevServer("http://localhost:3001"),
  kama.WithPath("static"),
)

// in gin

r.Use(k.Gin())

// in http

http.HandleFunc("/", k.Go())
```

> build or run with `-tags dev` to switch to dev mode, or it will be production mode by default.

# Tree

With option `kama.WithTree("/tree")`, kama add a new endpoint `/tree` that show files in the embed fs. It's helpful to check if the setting is fine. 

# License

[MIT](./LICENSE)
