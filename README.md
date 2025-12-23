# go-orders-api

Simple Go HTTP API using chi router.
# go-orders-api

Simple Go HTTP API using the `chi` router.

## Overview

This minimal project demonstrates a small HTTP server exposing a single endpoint.

## Prerequisites

- Go 1.18 or newer

## Run

Development:

```bash
go run main.go
```

Build and run:

```bash
go build -o bin/server .
./bin/server
```

The server listens on port `:8080` by default.

## Endpoints

- `GET /api` — returns `Hello World`

## Quick test

```bash
curl http://localhost:8080/api
# → Hello World
```

## Recent changes

- Middleware ordering fix: middlewares are now registered before routes in `main.go` to prevent the chi panic `all middlewares must be defined before routes on a mux`.
- Added several common middlewares in `main.go`: `Logger`, `Recoverer`, `RealIP`, and a 60s `Timeout`.

## Notes

- Uses `github.com/go-chi/chi/v5`.
- To change the listening port, edit the `Addr` field in `main.go` or wrap the server with an environment-configured value.

If you'd like, I can commit and push this change for you.
```

