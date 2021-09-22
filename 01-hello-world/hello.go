package main

import "fmt"

const french = "fr"
const spanish = "es"
const englishGreetingPrefix = "Hello, "
const frenchGreetingPrefix = "Bonjour, "
const spanishGreetingPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchGreetingPrefix
	case spanish:
		prefix = spanishGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("", ""))
}
