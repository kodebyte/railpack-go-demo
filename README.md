# railpack-go-demo

Minimal Dockerfile-free Go web server to test the Vessl **Railpack** build engine.

- `go.mod` + `main.go` only — no Dockerfile (so Railpack is selected, not Docker).
- Listens on `$PORT` (default 8080); `/` returns HTTP 200.
