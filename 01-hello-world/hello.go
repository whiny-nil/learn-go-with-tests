package main

import "fmt"

func Hello(greetee string) string {
	return "Hello, " + greetee
}

func main() {
	fmt.Println(Hello("world"))
}
