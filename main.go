package main

import (
	"fmt"
	"go-tdd/arraySlice"
	"go-tdd/integer"
)

const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", "English"))
	fmt.Println(integer.Add(2, 4))
	fmt.Println(arraySlice.SumAllTails([]int{1, 3, 5}))
}
