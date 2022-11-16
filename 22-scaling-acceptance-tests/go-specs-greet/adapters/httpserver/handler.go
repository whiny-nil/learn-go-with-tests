package httpserver

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/domain/interactions"
)

type GreetHandler struct{}

func (g GreetHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprint(res, go_specs_greet.Greet(name))
}

type CurseHandler struct{}

func (c CurseHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprint(res, go_specs_greet.Curse(name))
}
