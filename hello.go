package main

import "fmt"

const greetingPrefix = "Hello, "
const defaultGreeting = "World"

func PrintHello(name string) string {
	if name == "" {
		return greetingPrefix + defaultGreeting
	} else {
		return greetingPrefix + name
	}
}
func main() {
	fmt.Println(PrintHello("Malabika"))
}
