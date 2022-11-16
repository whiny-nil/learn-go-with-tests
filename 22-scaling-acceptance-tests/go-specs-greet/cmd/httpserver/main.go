package main

import (
	"net/http"

	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/adapters/httpserver"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/greet", httpserver.GreetHandler{})
	mux.Handle("/curse", httpserver.CurseHandler{})
	http.ListenAndServe(":8081", mux)
}
