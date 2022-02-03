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
	if language == spanish {
		return greetingPrefixSpanish + name
	} else if language == french {
		return greetingPrefixFrench + name
	}
	return greetingPrefixEnglish + name

}
func main() {
	fmt.Println(PrintHello("Malabika", ""))
}
