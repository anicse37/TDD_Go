package main

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

var (
	write = "Write"
	sleep = "Sleep"
)

func Sleeper(out io.Writer, value int) {
	fmt.Fprintln(out, sleep)
}
func Writer(out io.Writer, value int) {
	fmt.Fprintln(out, write)
}

func Countdown(out io.Writer, i int) {
	for i > 0 {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
		i--
	}
	fmt.Fprintf(out, "Go!")
}
func main() {
	buff := &bytes.Buffer{}
	Countdown(buff, 3)

}
func SleepCountdown(out io.Writer, i int) {
	for i > 0 {
		Writer(out, i)
		Sleeper(out, i)
		i--
	}
}
