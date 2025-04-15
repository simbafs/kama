# kama

> /kam-á/  
> 橘子的臺語  
> Tangerine in Taiwanese Hokkien

[中文版](./README-zh.md)

**Kama** is a lightweight Go package for building **frontend-backend separated web applications**.

In development mode, it proxies all unmatched HTTP requests to a frontend dev server (e.g. Vite, Next.js).  
In production mode, it serves static files from an embedded filesystem, with optional overrides from a local directory—no recompilation required.

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

kama.New(embeddedFS,
  kama.WithDevServer("http://localhost:3001"),
  kama.WithPath("static"),
)
```

# Usage

See the [\_example/](./_example/) directory for complete usage examples with both `net/http` and `gin`.

# License

[MIT](./LICENSE)
