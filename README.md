# kama

> kam-á 橘子的臺語  
> Tangerine in Taiwanese Hokkien

This package is designed for building applications with a separated frontend and backend using Go.
In development mode, all unhandled requests are redirected to a frontend development server (such as Next.js or Astro).
In production mode, all files in the specified directory are embedded into the output binary using Go's `embed` feature, allowing you to deploy the entire application as a single executable.

In practice, I supose you to write a Makefile or shell script to help you. In development mode, it start and watch frontend dev server and backend server. When you need to build in production mode, it builds frontend and move files to the specified directory in backend and run `go buil -tags prod`.

# overlay fs

The `static` directory in the current working directory will overlay the embedded filesystem, allowing you to overwrite any static file without recompiling.

# Example

See [example](./example/README.md)
