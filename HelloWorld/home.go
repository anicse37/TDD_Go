package main

import "fmt"

const constHelloString = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return constHelloString + name
}
func main() {
	fmt.Println(Hello("Aniket"))
}
