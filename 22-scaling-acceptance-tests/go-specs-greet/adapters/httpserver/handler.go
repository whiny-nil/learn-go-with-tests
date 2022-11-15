package httpserver

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/domain/interactions"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	name := req.URL.Query().Get("name")
	if path == "/greet" {
		fmt.Fprint(res, go_specs_greet.Greet(name))
	} else {
		fmt.Fprint(res, go_specs_greet.Curse(name))
	}
}
