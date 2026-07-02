package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Minimal Dockerfile-free Go web server for testing the Railpack build engine.
// Railpack detects go.mod -> `go build` -> runs the binary. Reads $PORT (Vessl
// injects it); "/" returns 200 so the health check + reachability probe pass.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Railpack Go Demo — v4-RECOVERED OK")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("railpack-go-demo listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
