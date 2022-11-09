package main

import (
	"log"
	"net/http"

	go_specs_greet "github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet"
)

func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatal(err)
	}
}
