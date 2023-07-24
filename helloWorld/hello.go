package main

import "fmt"

const (
	hindi  = "Hindi"
	french = "French"

	englishHelloPrefix = "Hello, "
	hindiHelloPrefix   = "Namaste, "
	frenchHelloPrefix  = "Bonjour, "
)

// Prints hello based on provided name and language
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(lang)

	return prefix + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case hindi:
		prefix = hindiHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
