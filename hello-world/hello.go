package hello

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
