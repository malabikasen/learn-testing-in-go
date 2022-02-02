package main

import "fmt"

const greetingPrefix = "Hello, "

func PrintHello(name string) string {
	return greetingPrefix + name
}
func main() {
	fmt.Println(PrintHello("Malabika"))
}
