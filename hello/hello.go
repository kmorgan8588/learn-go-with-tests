package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello takes a string and appends it to a prefix greeting
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world"))
}
