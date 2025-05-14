package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	spanish            = "Spanish"
)

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}
