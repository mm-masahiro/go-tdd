package main

import "fmt"

const prefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
