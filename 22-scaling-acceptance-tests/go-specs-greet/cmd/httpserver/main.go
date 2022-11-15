package main

import (
	"net/http"

	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	http.ListenAndServe(":8081", handler)
}
