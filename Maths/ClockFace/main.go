package main

import (
	"os"
	"time"

	clockFace "github.com/anicse37/TDD_Go/Maths"
)

/*-----------------------------------------------------------------------------------------------------*/
func main() {

	t := time.Now()
	clockFace.SVGWritter(os.Stdout, t)

}
