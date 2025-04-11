package main

import (
	"bytes"
	"fmt"
	"io"
)

func Greet(write io.Writer, name string) {
	fmt.Fprintf(write, "Hello, %s", name)
}
func main() {
	buff := bytes.Buffer{}
	Greet(&buff, "Slim Shady")
	fmt.Println(buff.String())
}
