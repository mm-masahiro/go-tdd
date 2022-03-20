package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix

	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
