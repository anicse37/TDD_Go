package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Aniket"
	fmt.Fprintf(os.Stdout, "My name is %s\n", name)
}
