package main

import (
	"os"
	"time"

	clockFace "github.com/anicse37/TDD_Go/Maths"
)

/*-----------------------------------------------------------------------------------------------------*/
func main() {
	for {
		f, _ := os.Create("clock.svg")
		t := time.Now()
		clockFace.SVGWritter(f, t)
		f.Close()
		time.Sleep(1 * time.Second)
	}
}
