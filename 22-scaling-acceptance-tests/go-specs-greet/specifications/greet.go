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

type MeanGreeter interface {
	Curse(string) (string, error)
}

type CurseAdapter func(name string) string

func (c CurseAdapter) Curse(name string) (string, error) {
	return c(name), nil
}

func CurseSpecification(t testing.TB, meany MeanGreeter) {
	got, err := meany.Curse("Marc")
	assert.NoError(t, err)
	assert.Equal(t, got, "I hate you, Marc")
}
