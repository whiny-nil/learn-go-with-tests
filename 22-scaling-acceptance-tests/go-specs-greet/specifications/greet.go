package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(string) (string, error)
}

type GreetAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Marc")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Marc")
}
