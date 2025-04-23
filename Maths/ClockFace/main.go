package main

import (
	"os"
	"time"

	clockFace "aniket/Maths"
)

/*-----------------------------------------------------------------------------------------------------*/
func main() {

	t := time.Now()
	clockFace.SVGWritter(os.Stdout, t)

}
