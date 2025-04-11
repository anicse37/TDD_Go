package dependency

import (
	"fmt"
	"io"
)

type Greeter struct {
	Writer io.Writer
}

func (g *Greeter) Greet(name string) {
	fmt.Fprintf(g.Writer, "Hello, %s", name)
}
