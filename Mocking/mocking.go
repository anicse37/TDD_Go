package main

import (
	"fmt"
	"io"
	"time"
)

// Countdown prints a countdown from 3 to out.
func Countdown(out io.Writer, i int) {

	for i > 0 {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
		i--
	}
	fmt.Fprintf(out, "Go!")
}
