package go_specs_greet

import (
	"fmt"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprintf(res, Greet(name))
}
