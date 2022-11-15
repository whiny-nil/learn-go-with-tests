package go_specs_greet_test

import (
	"testing"

	go_specs_greet "github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/domain/interactions"
	"github.com/whiny-nil/learn-go-with-tests/22-scaling-acceptance-tests/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(go_specs_greet.Greet))
}

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(t, specifications.CurseAdapter(go_specs_greet.Curse))
}
