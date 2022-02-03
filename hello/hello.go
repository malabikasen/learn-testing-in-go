package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const greetingPrefixEnglish = "Hello, "
const greetingPrefixSpanish = "Hola, "
const greetingPrefixFrench = "Bonjour, "

// greetingPrefix is a private function
// hence name starts with lowercase letters
// we do not want the internals of the algorithm to be exposed hence
// Public functions are uppercase
// prefix is named return value
func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = greetingPrefixSpanish
	case french:
		prefix = greetingPrefixFrench
	default:
		prefix = greetingPrefixEnglish
	}

	return
}

func PrintHello(name string, language string) string {

	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
func main() {
	fmt.Println(PrintHello("Malabika", ""))
}
