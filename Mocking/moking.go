package main

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

// Countdown prints a countdown from 3 to out.
func Countdown(out io.Writer, i int) {
	fmt.Fprintf(out, "%v", i)
}
func main() {
	i := 3
	fmt.Println("hello")
	buff := bytes.Buffer{}
	for i != 0 {
		time.Sleep(1 * time.Second)
		Countdown(&buff, i)
		fmt.Print(&buff, "\n")
		i--
	}
	fmt.Println("hello")
}
