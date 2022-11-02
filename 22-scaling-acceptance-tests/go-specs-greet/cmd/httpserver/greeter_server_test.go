package main_test

import (
	"testing"

	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/drivers"
	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/specifications"
)

func TestGreeterService(t *testing.T) {
	driver := drivers.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
