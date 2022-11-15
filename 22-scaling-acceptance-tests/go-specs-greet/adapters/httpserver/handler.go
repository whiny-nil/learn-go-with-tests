package httpserver

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/domain/interactions"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprintf(res, go_specs_greet.Greet(name))
}
