package go_specs_greet

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func Curse(name string) string {
	return fmt.Sprintf("I hate you, %s", name)
}
