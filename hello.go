package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const greetingPrefixEnglish = "Hello, "
const greetingPrefixSpanish = "Hola, "
const greetingPrefixFrench = "Bonjour, "

func PrintHello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	var prefix string

	switch language {
	case spanish:
		prefix = greetingPrefixSpanish
	case french:
		prefix = greetingPrefixFrench
	default:
		prefix = greetingPrefixEnglish
	}

	return prefix + name

}
func main() {
	fmt.Println(PrintHello("Malabika", ""))
}
